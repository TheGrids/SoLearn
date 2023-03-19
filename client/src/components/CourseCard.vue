<template>
  <Card @click="redirect" class="course item">
    <template #header>
      <div style="overflow: hidden">
        <Button @click.stop="sendLike" rounded class="like">
          <i class="pi"
             :class="{'pi-heart-fill': likeSend, 'pi-heart': !likeSend}"
             style="font-size: 1rem;"
          ></i>
        </Button>
        <img
            style="height: 100px"
            :src="course.image"
            alt="Course preview"
        />
      </div>
    </template>
    <template #title>
      <div style="line-height: 1">
        <span class="title">{{ course.name }}</span>
      </div>
    </template>
    <template #content style="overflow: hidden">
      <div style="overflow: hidden; padding: 0; height: min-content">
        <div style="text-overflow: ellipsis">
          {{ description }}
        </div>
      </div>
    </template>
  </Card>
</template>

<script>
import Card from 'primevue/card';

export default {
  name: "CourseCard",
  components: {
    Card,
  },
  props: {
    course: Object,
  },
  data() {
    return {
      likeSend: false,
    };
  },
  computed: {
    description() {
      if (this.course.description.length <= 100) {
        return this.course.description;
      } else {
        return this.course.description.substring(0, 100);
      }
    },
  },
  methods: {
    redirect() {
      this.$router.push(`/courses/${this.course.id}`)
    },
    sendLike() {
      this.likeSend = !this.likeSend;
      if (this.likeSend) {
        this.$toast.add({severity: 'info', summary: 'Добавлено в избранное', life: 10000});
      } else {
        this.$toast.add({severity: 'info', summary: 'Убрано из избранного', life: 10000});
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.course {
  position: relative;
  font-family: 'Nunito', sans-serif;
  border-radius: 15px;
  overflow: hidden;
  border: 3px solid transparent;
  transition: 250ms;

  .title {
    line-height: 1;
    font-weight: 700;
    font-size: 20px;
  }

  img {
    object-fit: cover;
    transition: 250ms;
  }


  &:hover {
    cursor: pointer;
    border: 3px solid #ffffff;
    box-shadow: 0 0 5px 5px #EDEDED;


    .title {
      color: #4911c0;
    }


    img {
      transform: scale(1.05);
    }
  }
}

.item {
  grid-column: span 1;
}

.like {
  width: 30px;
  min-height: 30px;
  display: flex;
  justify-content: center;
  align-items: center;
  position: absolute;
  background: #333;
  border: none;
  top: 10px;
  right: 10px;
  z-index: 10;
}
</style>