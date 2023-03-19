<template>
  <div style="padding: 15px">
    <header class="app-header">
      <div class="app-header__content">
        <router-link to="/" class="title">
          <span style="color: #383B41;">So</span><span style="color: #697187">Learn</span>
        </router-link>
        <div class="app-header-menu">
          <router-link v-if="isAuthorized && $route.path !== '/profile' && $route.path !== '/courses'" to="/profile">
            <Button icon="pi pi-user" link style="margin-right: 15px" v-tooltip.bottom="'Профиль'"/>
          </router-link>
          <Button
              v-if="authStore.user"
              @click="signOutClicked"
              :loading="signOutLoading"
              v-tooltip.bottom="'Выйти'"
              rounded
              class="logout"
              icon="pi pi-sign-out"
              link
          />
        </div>
      </div>
    </header>
    <main class="main-content">
      <RouterView/>
    </main>
  </div>
</template>

<script>
import {mapStores} from "pinia";
import {useAuthStore} from "@/stores/auth";

export default {
  name: "MainLayout",
  data() {
    return {
      signOutLoading: false,
    };
  },
  computed: {
    ...mapStores(useAuthStore),
    isAuthorized() {
      return !!this.authStore.user;
    },
  },
  methods: {
    async signOutClicked() {
      this.signOutLoading = true;
      await this.authStore.logout();
      this.$router.replace({name: 'login'});
      this.signOutLoading = false;
    }
  }
}
</script>

<style lang="scss" scoped>
.main-content {
  width: 100%;
  margin: 0 auto;
  max-width: 1440px;
}

.app-header {
  display: flex;
  flex-direction: row;
  width: 100%;
  justify-content: center;

  .app-header__content {
    display: flex;
    flex-direction: row;
    width: 100%;
    max-width: 1440px;
    padding: 20px 0;
    justify-content: space-between;
    align-items: center;
  }

  .title {
    span {
      font-size: 32px;
      font-weight: 700;
    }
  }
}

.logout {
  min-width: 30px;
  min-height: 30px;
  display: flex;
  justify-content: center;
  align-center: center;
}

.app-header-menu {
  display: flex;
  flex-direction: row;
}
</style>