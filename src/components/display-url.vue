<!-- eslint-disable vue/no-v-html -->
<template>
  <div class="card border-success-subtle mb-3">
    <!-- Celebration banner -->
    <div class="ots-success-banner">
      <span class="ots-success-icon"><i class="fas fa-circle-check" /></span>
      <p class="ots-success-title" v-html="$t('title-secret-created')" />
    </div>
    <div class="card-body">
      <p class="text-muted mb-3" v-html="$t('text-pre-url')" />
      <div class="input-group mb-3">
        <input
          ref="secretUrl"
          class="form-control font-monospace"
          type="text"
          readonly
          :value="secretUrl"
          @focus="$refs.secretUrl.select()"
        >
        <app-clipboard-button
          :content="secretUrl"
          :title="$t('tooltip-copy-to-clipboard')"
        />
      </div>
      <p class="text-muted mb-1" v-html="$t('text-burn-hint')" />
      <p
        v-if="expiresAt"
        class="mb-0 expiry-notice"
      >
        <i class="fas fa-clock me-1" />
        {{ $t('text-burn-time') }}
        <strong>{{ expiresAt.toLocaleString() }}</strong>
      </p>
    </div>
  </div>
</template>
<style scoped>
.expiry-notice {
  font-size: .875rem;
  color: #2e7d32;
  font-weight: 700;
}
[data-bs-theme=dark] .expiry-notice {
  color: #86efac;
}
</style>
<script>
import appClipboardButton from './clipboard-button.vue'
import appQrButton from './qr-button.vue'

export default {
  components: { appClipboardButton, appQrButton },
  computed: {
    secretUrl() {
      return [
        window.location.href.split('#')[0],
        encodeURIComponent([
          this.secretId,
          this.securePassword,
        ].join('|')),
      ].join('#')
    },
  },

  data() {
    return {
      popover: null,
    }
  },

  mounted() {
    // Give the interface a moment to transistion and focus
    window.setTimeout(() => this.$refs.secretUrl.focus(), 100)
  },

  name: 'AppDisplayURL',
  props: {
    expiresAt: {
      default: null,
      required: false,
      type: Date,
    },

    secretId: {
      required: true,
      type: String,
    },

    securePassword: {
      required: true,
      type: String,
    },
  },
}
</script>
