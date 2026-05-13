<template>
  <nav class="navbar navbar-expand-lg">
    <div class="container">
      <a
        class="navbar-brand d-flex align-items-center gap-2"
        href="#"
        @click.prevent="$root.navigate('/')"
      >
        <img
          class="logo"
          src="/favicon.png"
          alt="logo"
        >
        <span v-if="!$root.customize.disableAppTitle" class="brand-text">Webman OTS</span>
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
        <ul class="navbar-nav ms-auto mb-2 mb-lg-0 me-3 align-items-center">
          <li class="nav-item">
            <a
              class="nav-link nav-link-new"
              href="#"
              @click.prevent="$root.navigate('/')"
            >
              <i class="fas fa-plus me-1" /> {{ $t('btn-new-secret') }}
            </a>
          </li>
        </ul>
        <form
          v-if="!$root.customize.disableThemeSwitcher"
          class="d-flex align-items-center gap-2"
        >
          <i class="fas fa-sun theme-icon" />
          <div class="form-check form-switch mb-0">
            <input
              id="themeswitch"
              v-model="$root.darkTheme"
              class="form-check-input"
              type="checkbox"
              role="switch"
            >
          </div>
          <i class="fas fa-moon theme-icon" />
        </form>
      </div>
    </div>
  </nav>
</template>

<style scoped>
.logo {
  width: 30px;
  height: 30px;
  object-fit: contain;
  filter: drop-shadow(0 1px 3px rgba(0,0,0,.3));
}

.form-check-input {
  margin: 0 !important;
}

.brand-text {
  font-weight: 700;
  font-size: 1.1rem;
  letter-spacing: .02em;
  color: white;
}

.nav-link-new {
  background: rgba(255,255,255,.15);
  border-radius: 7px;
  transition: background .18s ease;
}

.nav-link-new:hover {
  background: rgba(255,255,255,.25) !important;
}

.theme-icon {
  color: rgba(255,255,255,.75);
  font-size: .88rem;
}
</style>

<script>
export default {
  name: 'AppNavbar',
}
</script>
