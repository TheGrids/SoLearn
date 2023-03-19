<template>
  <div style="min-height: 100vh; display: flex; flex-direction: column; justify-content: space-between">
    <div class="font-bold text-4xl mt-4 mb-14" style="color: #232323">
      Доступные курсы
    </div>
    <div class="wrapper">
      <CourseCard
          v-for="course in courses"
          :course="course"
          :key="course.id"
      />
    </div>
    <footer>
      <div class="container">
        <div class="footer">
          <div class="footer-logo">
            <span style="font-weight: bold;color: #383B41;">So</span><span style="font-weight: bold;color: #697187">Learn</span>
          </div>
          <div class="footer-info">
            <div class="footer-info__item">
              <div class="item-name">Сообщество</div>
              <div class="item-list">github</div>
              <div class="item-list">vk</div>
              <div class="item-list">github</div>
            </div>

            <div class="footer-info__item">
              <div class="item-name">Сервис</div>
              <div class="item-list">О нас</div>
              <div class="item-list">Новости</div>
              <div class="item-list">Вакансии</div>
            </div>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<script>
import CourseCard from "@/components/CourseCard.vue";
import {api} from "@/http";

export default {
  name: "AvailableCoursesView",
  components: {CourseCard},
  mounted() {
    this.getAllCourses();
  },
  methods: {
    async getAllCourses() {
      const response = await api.get('getcourses');
      this.courses = response.data.data;
    },
  },
  data() {
    return {
      courses: [],
    }
  }
}
</script>

<style lang="scss" scoped>
.wrapper {
  width: 100%;
  display: grid;
  grid-template-columns: 1fr;
  grid-gap: 20px;

  @media (min-width: 768px) {
    grid-template-columns: 1fr 1fr;
  }

  @media (min-width: 1000px) {
    grid-template-columns: 1fr 1fr 1fr;
  }

  @media (min-width: 1400px) {
    grid-template-columns: 1fr 1fr 1fr 1fr;
  }
}

.footer {
  padding: 50px 20px;
  width: 100%;
  color: #383B41;
  display: flex;
  justify-content: space-between;

  .footer-logo {
    font-size: 32px;
  }

  .footer-info {
    display: flex;

    .footer-info__item {
      margin-left: 200px;

      .item-name {
        font-size: 24px;
        font-weight: 700;
      }

      .item-list {
        font-size: 20px;
      }
    }
  }
}
</style>