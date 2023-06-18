<!-- eslint-disable -->
<template>
    <swiper
      class="swiper"
      :modules="modules"
      :space-between="30"
      :slides-per-view="2"
      :free-mode="true"
      :pagination="{ clickable: true }"
     
    >
      <swiper-slide class="slide" v-for="image in partImages.partPhotos" :key="image.partPhoto">
        <img class="swiper-image" :src="'data:image/png;base64,'+ image.partPhoto " alt="test">
      </swiper-slide>
    </swiper>
</template>
  
  <script>
    import { defineComponent } from 'vue'
    import { Pagination, FreeMode } from 'swiper'
    import { Swiper, SwiperSlide } from 'swiper/vue'
    import 'swiper/css'
    import 'swiper/css/pagination'
    import 'swiper/css/free-mode'
    export default defineComponent({
      name: 'swiper-example-free-mode',
      title: 'Free mode / no fixed positions',
      url: import.meta.url,
      components: {
        Swiper,
        SwiperSlide
      },
      props: {
        partImages:{
            type: Object,
            default() {
                return {}
            }
        }
       },
       methods : {
        getinfo () {
          console.log(this.partImages.partPhotos)
        }
       },
      setup() {
        return {
          modules: [Pagination, FreeMode]
        }
      }
    })
  </script>
  
  <style lang="scss" scoped>
    @import '@/styles/variables.scss';
    @import '@/styles/mixins.scss';
    @import './style.scss';
   
    .swiper {
      @include swiper-wrapper();
      height: 100%;
      width: 100%;
    }
    .slide {
      @include swiper-slide();
    }
  
    .swiper-image{
      width: 100%;
      height: 250px;
    }
  </style> 