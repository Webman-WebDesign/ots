<template>
  <nav class="navbar navbar-expand-lg bg-primary-subtle">
    <div class="container-fluid">
      <a
        class="navbar-brand"
        href="#"
        @click.prevent="$root.navigate('/')"
      >
        <img
          class="mr-1 logo"
          src="/favicon.png"
        >
        <span v-if="!$root.customize.disableAppTitle">Webman - One Time Secrets</span>
      </a>

      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarSupportedContent"
        aria-controls="navbarSupportedContent"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon" />
      </button>

      <div
        id="navbarSupportedContent"
        class="collapse navbar-collapse"
      >
        <ul class="navbar-nav ms-auto mb-2 mb-lg-0 me-2">
          <li class="nav-item">
            <a
              class="nav-link"
              href="https://webman-webdesign.de"
            >
              <i class="fas fa-circle-info" /> Webman-Webdesign
            </a>
          </li>
          <li class="nav-item">
            <a
              class="nav-link"
              href="#"
              @click.prevent="$root.navigate('/')"
            >
              <i class="fas fa-plus" /> {{ $t('btn-new-secret') }}
            </a>
          </li>
        </ul>
        <form
          v-if="!$root.customize.disableThemeSwitcher"
          class="d-flex align-items-center"
        >
          <i class="fas fa-sun me-2" />
          <div class="form-check form-switch">
            <input
              id="themeswitch"
              v-model="$root.darkTheme"
              class="form-check-input"
              type="checkbox"
              role="switch"
            >
          </div>
          <i class="fas fa-moon" />
        </form>
      </div>
    </div>
  </nav>
</template>

<style scoped>
.logo {
  width: 28px;
}
</style>

<script>
export default {
  name: 'AppNavbar',
}
</script>
