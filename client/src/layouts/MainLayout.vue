<template>
  <div style="padding: 15px">
    <header class="app-header">
      <div class="app-header__content">
        <p class="title">
          <span style="color: #383B41;">So</span><span style="color: #697187">Learn</span>
        </p>
        <Button rounded link>
          <i class="pi pi-sign-out" size="small" style="font-size: 1.5rem;"></i>
        </Button>
      </div>
    </header>
    <main class="main-content">
      <RouterView/>
    </main>
  </div>
</template>

<script>
export default {
  name: "MainLayout"
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
</style>