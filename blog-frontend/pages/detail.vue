<template>
    <Navbar />
    <!-- Content -->
    <div class="container mt-4">
      <div class="row">
        <!-- Main Content -->
        <div class="col-lg-8">
          <!-- Breadcrumb -->
          <nav aria-label="breadcrumb">
            <ol class="breadcrumb">
              <li class="breadcrumb-item"><NuxtLink to="/">Home</NuxtLink></li>
              <li class="breadcrumb-item"><NuxtLink to="/">HTML</NuxtLink></li>
              <li class="breadcrumb-item active" aria-current="page">Apa Itu HTML?</li>
            </ol>
          </nav>

          <!-- Article -->
          <article>
            <h2>Mengenal Apa Itu HTML?</h2>
            <p class="text-muted">By Ridho Satriawan - 22 Agustus 2022</p>

            <h4>Apa itu HTML?</h4>
            <p>
              Buat kamu yang pernah membuat website pastinya sudah tidak asing lagi dengan yang namanya HTML. 
              Sebenarnya apa sih HTML itu? Kenapa menggunakan HTML untuk membuat website?
            </p>

            <h4>Pengertian</h4>
            <p>
              HTML merupakan singkatan dari HyperText Markup Language, merupakan bahasa markup standar yang digunakan 
              untuk menampilkan dokumen ke Web Browser seperti Google Chrome dan Edge. HTML biasanya digunakan bersama 
              teknologi lain seperti CSS dan JavaScript.
            </p>

            <h4>Sejarah Singkat</h4>
            <p>
              HTML pertama kali dicetus pada tahun 1980 oleh Tim Berners-Lee, pada saat itu dia sedang bekerja di CERN, 
              mengusulkan sistem yang berfungsi untuk mengelola dokumen perusahaan.
            </p>

            <!-- Example Code Block -->
            <pre>
              <code>
                &lt;html&gt;
                  &lt;head&gt;
                    &lt;title&gt;Example&lt;/title&gt;
                  &lt;/head&gt;
                  &lt;body&gt;
                    &lt;h1&gt;Hello, World!&lt;/h1&gt;
                  &lt;/body&gt;
                &lt;/html&gt;
              </code>
            </pre>

            <h4>Penutup</h4>
            <p>
              HTML terus berkembang hingga saat ini dan tetap menjadi dasar dari pengembangan website.
            </p>
          </article>
        </div>

        <!-- Sidebar -->
        <Sidebar />
      </div>
    </div>
    <Footer />
</template>

<style scoped>
.breadcrumb {
  background-color: transparent;
  padding: 0;
  margin-bottom: 1rem;
}
.breadcrumb-item a {
  text-decoration: none;
  color: #00A885;
}
.breadcrumb-item.active {
  color: #6c757d;
}
pre {
  background-color: #f8f9fa;
  padding: 15px;
  border-radius: 5px;
  font-size: 0.9rem;
}
</style>