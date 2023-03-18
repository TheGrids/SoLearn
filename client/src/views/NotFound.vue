<template>
  <div class="not-found-page">
    <span class="title">404</span>
  </div>
</template>

<script>
export default {
  name: "NotFound"
}
</script>

<style lang="scss" scoped>
.not-found-page {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;

  .title {
    font-size: 20vw;
    font-weight: 700;
  }
}
</style>