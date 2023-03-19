<template>
  <div style="margin: 20px 0">
    <p class="course-creation-title">Создание курса</p>
    <p class="course-creation-title mt-10">Общее</p>

    <div class="row" style="gap: 10px; width: 100%">
      <div class="column" style="width: 100%">
        <span class="text-2xl mb-2">Название</span>
        <InputText placeholder="Введите название"/>
      </div>

      <div class="column" style="width: 100%">
        <span class="text-2xl mb-2">Категория</span>
        <InputText placeholder="Введите категорию"/>
      </div>
    </div>
    <div class="column mt-5">
      <span class="text-2xl mb-2">Описание</span>
      <InputText placeholder="Введите описание"/>
    </div>
    <div class="mt-10 text-2xl mb-10">Программа курса</div>
    <TransitionGroup name="list" tag="div">
      <CourseTheme
          v-for="(item, idx) in themes"
          :theme="item"
          @remove="removeTheme"
          @update:theme="updateTheme"
          :step-number="++idx"
          :key="item.id"
      />
    </TransitionGroup>
    <div
        class="column mt-4"
        style="gap: 15px; max-width: 200px"
    >
      <Button class="app-button" @click="addTheme" label="Добавить тему"/>
      <Button @click="createCourse" class="app-button" label="Сохранить"/>
    </div>
  </div>
</template>

<script>
import CourseTheme from "@/components/CourseTheme.vue";
import {ref} from "vue";
import {v4 as uuidV4} from 'uuid';

export default {
  name: "CreateCourseView",
  components: {CourseTheme},
  setup() {
    const themes = ref([]);

    return {
      themes,
    }
  },
  methods: {
    addTheme() {
      this.themes.push({
        id: uuidV4(),
        title: '',
        subThemes: [],
      });
    },
    updateTheme(updatedTheme) {
      this.themes.map((theme) => {
        if (theme.id === updatedTheme.id) {
          theme = updatedTheme;
        }
        return theme;
      });
    },
    removeTheme(id) {
      this.themes = this.themes.filter((theme) => theme.id !== id);
    },
    createCourse() {
      console.log(this.themes)
    },
  }
}
</script>

<style lang="scss" scoped>
.course-creation-title {
  font-family: 'Nunito', sans-serif;
  font-style: normal;
  font-weight: 700;
  font-size: 32px;
  line-height: 44px;
  margin-bottom: 20px;
}

.list-move,
.list-enter-active,
.list-leave-active {
  transition: all 0.5s ease;
}

.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: translateY(7px);
}

.list-leave-active {
  position: absolute;
}

.app-button {
  border-radius: 25px;
  background: #181818;
  color: #f2f2f2;
}

.p-dialog-mask {
  z-index: 1 !important;
}
</style>