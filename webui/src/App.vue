<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {}
</script>

<template>
  <div>
    <header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
      <a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">WASAText</a>
      <button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon" />
      </button>
    </header>

    <div class="container-fluid">
      <div class="row">
        <nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
          <div class="position-sticky pt-3 sidebar-sticky">
            <h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
              <span>General</span>
            </h6>
            <ul class="nav flex-column">
              <li class="nav-item">
                <RouterLink to="/chats" class="nav-link">
                  <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home" /></svg>
                  Chats
                </RouterLink>
              </li>
              <li class="nav-item">
                <RouterLink to="/profile" class="nav-link">
                  <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#layout" /></svg>
                  Profile Settings
                </RouterLink>
              </li>
              <li class="nav-item">
                <RouterLink to="/" class="nav-link">
                  <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#key" /></svg>
                  Login
                </RouterLink>
              </li>
            </ul>
          </div>
        </nav>

        <main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
          <RouterView />
        </main>
      </div>
    </div>
  </div>
</template>

<style>
</style>
