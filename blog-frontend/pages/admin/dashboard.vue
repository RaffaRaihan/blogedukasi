<template>
  <div class="container-fluid">
    <div class="row">
      <!-- Sidebar -->
      <AdminSidebar />

      <!-- Main Content -->
      <main class="col-md-9 col-lg-10 p-4">
        <h2>Workspace admins</h2>
        <!-- Filters -->
        <div class="d-flex align-items-center mb-4 mt-4">
          <input type="text" class="form-control me-2" placeholder="Search" />
          <select class="form-select me-2">
            <option>All</option>
            <option>Admin</option>
            <option>User</option>
          </select>
        </div>
        <!-- Admin Cards -->
        <div class="mb-4">
          <div class="d-flex">
            <div class="card me-2" style="width: 150px;">
              <div class="card-body text-center">
                <h6 class="card-title">John Johnson</h6>
                <p class="text-muted">CEO</p>
                <button class="btn btn-sm btn-danger">Remove</button>
              </div>
            </div>
            <div class="card d-flex align-items-center justify-content-center" style="width: 150px; border: dashed 2px #ddd;">
              <button class="btn btn-outline-primary">+ Add Admin</button>
            </div>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

  