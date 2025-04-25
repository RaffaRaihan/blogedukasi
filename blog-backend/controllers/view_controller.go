package controllers

import (
	"blog-backend/config"
	"blog-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TrackView(c *gin.Context) {
	var input struct {
		ArticleID uint `json:"article_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Optional: Validasi bahwa artikelnya ada
	var article models.Article
	if err := config.GetDB().First(&article, input.ArticleID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artikel tidak ditemukan"})
		return
	}

	view := models.ArticleView{
		ArticleID: input.ArticleID,
		ViewedAt:  time.Now(),
	}

	if err := config.GetDB().Create(&view).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan view"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "View berhasil dicatat"})
}

// GetPopularArticles mengambil daftar artikel yang paling sering dilihat
func GetPopularArticles(c *gin.Context) {
	var results []struct {
		ArticleID uint   `json:"article_id"`
		Title     string `json:"title"`
		CreatedAt string `json:"created_at"`
		Views     int64  `json:"views"`
	}

	// Query dengan JOIN ke tabel articles
	config.GetDB().Table("article_views").
		Select("articles.id as article_id, articles.title, articles.created_at, COUNT(article_views.article_id) as views").
		Joins("JOIN articles ON articles.id = article_views.article_id").
		Group("article_views.article_id, articles.id, articles.title, articles.created_at").
		Order("views DESC").
		Limit(10).
		Scan(&results)

	// Kirim response ke frontend
	c.JSON(http.StatusOK, results)
}

func GetArticleStats(c *gin.Context) {
    var stats []models.ArticleStatus
    config.GetDB().Model(&models.Article{}).
        Select("status, COUNT(*) as count").
        Group("status").
        Find(&stats)

    c.JSON(http.StatusOK, stats)
}

func GetWeeklyArticleStats(c *gin.Context) {
	type Result struct {
		Day   int
		Count int
	}

	// Ambil query param
	week := c.Query("week")
	month := c.Query("month")
	year := c.Query("year")

	// Base query
	query := `
		SELECT EXTRACT(DOW FROM created_at) AS day, COUNT(*) AS count
		FROM articles
		WHERE 1=1
	`

	// Dynamic filter
	var params []interface{}
	if week != "" {
		query += " AND EXTRACT(WEEK FROM created_at) = ?"
		params = append(params, week)
	}
	if month != "" {
		query += " AND EXTRACT(MONTH FROM created_at) = ?"
		params = append(params, month)
	}
	if year != "" {
		query += " AND EXTRACT(YEAR FROM created_at) = ?"
		params = append(params, year)
	}

	query += " GROUP BY day"

	db := config.GetDB()
	var dataDB []Result
	db.Raw(query, params...).Scan(&dataDB)

	// Inisialisasi struktur output
	labels := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	data := make([]int, 7)

	// Mapping berdasarkan PostgreSQL EXTRACT(DOW)
	for _, r := range dataDB {
		index := (r.Day + 6) % 7 // Shift agar Monday = index 0
		data[index] = r.Count
	}

	c.JSON(http.StatusOK, gin.H{
		"labels": labels,
		"data":   data,
	})
}


func GetViewsPerCategory(c *gin.Context) {
	type Result struct {
		Category string
		Total    int
	}

	var results []Result

	db := config.GetDB()
	db.
		Raw(`
			SELECT c.name AS category, COUNT(v.id) AS total
			FROM article_views v
			JOIN articles a ON v.article_id = a.id
			JOIN categories c ON a.category_id = c.id
			GROUP BY c.name
		`).
		Scan(&results)

	var labels []string
	var data []int

	for _, r := range results {
		labels = append(labels, r.Category)
		data = append(data, r.Total)
	}

	c.JSON(http.StatusOK, gin.H{
		"labels": labels,
		"data":   data,
	})
}

func GetTopAuthorsByArticles(c *gin.Context) {
	type Result struct {
		Author string
		Total  int
	}

	var results []Result

	db := config.GetDB()
	db.
		Raw(`
			SELECT u.name AS author, COUNT(a.id) AS total
			FROM articles a
			JOIN users u ON a.author_id = u.id
			GROUP BY u.name
			ORDER BY total DESC
		`).
		Scan(&results)

	var labels []string
	var data []int

	for _, r := range results {
		labels = append(labels, r.Author)
		data = append(data, r.Total)
	}

	c.JSON(http.StatusOK, gin.H{
		"labels": labels,
		"data":   data,
	})
}