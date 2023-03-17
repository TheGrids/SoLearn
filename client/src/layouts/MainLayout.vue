<template>
  <header class="app-header">
    <div class="app-header__content">
      <p class="title">
        <span style="color: #383B41;">So</span><span style="color: #697187">Learn</span>
      </p>
      <Button rounded/>
    </div>
  </header>
  <main>
    <RouterView/>
  </main>
</template>

<script>
export default {
  name: "MainLayout"
}
</script>

<style lang="scss" scoped>
.app-header {
  display: flex;
  flex-direction: row;
  width: 100vw;
  justify-content: center;

  .app-header__content {
    display: flex;
    flex-direction: row;
    width: 100%;
    max-width: 1440px;
    padding: 20px;
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