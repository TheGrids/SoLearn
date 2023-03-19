<template>
  <Transition name="fade">
    <div v-if="course && !pageLoading" class="container" style="font-size: 40px">
      <h1 class="title mb-3">
        {{ course.name }}
      </h1>
      <p class="text-lg description">
        {{ course.description }}
      </p>
    </div>
    <div v-else-if="pageLoading" class="flex center loader">
      <ProgressSpinner
          style="width: 100px; height: 100px;"
          strokeWidth="8" fill="var(--surface-ground)"
          animationDuration=".5s"
      />
    </div>
    <div v-else-if="!pageLoading && !course" class="flex center" style="align-items: center">
      Курс не найден
    </div>
  </Transition>
</template>

<script>
import ProgressSpinner from "primevue/progressspinner";
import {api} from "@/http";

export default {
  name: "CourseView",
  components: {
    ProgressSpinner,
  },
  data() {
    return {
      course: null,
      pageLoading: true,
    }
  },
  mounted() {
    this.getCourseById();
  },
  methods: {
    async getCourseById() {
      const id = this.$route.params.id;
      try {
        const response = await api.get(`getcourse/${id}`);
        this.course = response.data.data;
      } catch (error) {
        console.log(error.response)
        this.showError(error.response.data);
      } finally {
        this.pageLoading = false;
      }
    },
    showError(message) {
      this.$toast.add({severity: 'error', summary: message, life: 10000});
    },
  }
}
</script>

<style lang="scss" scoped>
.container {
  background: #EDEDED;
  border-radius: 20px;
  width: 100%;
  padding: 10px 25px;
}

.title {
  font-weight: 700;
  font-size: 48px;
  line-height: 65px;
  color: #383B41;
}

.description {
  font-weight: 400;
  line-height: 1.2;
  font-size: 32px;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.loader {
  position: absolute;
  width: 100%;
  align-items: center;
}
</style>