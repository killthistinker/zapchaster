 <template>
  <swiper
    class="swiper"
    :modules="modules"
    :space-between="30"
    :slides-per-view="3"
    :free-mode="true"
    :pagination="{ clickable: true }"
    :autoplay="{ delay: 1000 }"
    :loop="true"
    effect="coverflow"
    :speed=speedValue
  >
    <swiper-slide class="slide"><img class="swiper-image" src="../../../public/hyndai.jpg" alt="image"></swiper-slide>
    <swiper-slide class="slide"><img class="swiper-image" src="../../../public/kia.jpg" alt="image"></swiper-slide>
    <swiper-slide class="slide"><img class="swiper-image" src="../../../public/hyndai.jpg" alt="image"></swiper-slide>
    <swiper-slide class="slide"><img class="swiper-image" src="../../../public/kia.jpg" alt="image"></swiper-slide>
    <swiper-slide class="slide"><img class="swiper-image" src="../../../public/hyndai.jpg" alt="image"></swiper-slide>
  </swiper>
</template>

<script>
  import { defineComponent } from 'vue'
  import { Pagination, FreeMode } from 'swiper'
  import { Swiper, SwiperSlide } from 'swiper/vue'
  import { EffectCoverflow } from 'swiper';
  import SwiperCore, { Autoplay } from 'swiper/core';
import 'swiper/swiper-bundle.css';

SwiperCore.use([Autoplay]);
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
    setup() {
      return {
        modules: [Pagination, FreeMode, EffectCoverflow],
        speedValue: 1500
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

@media screen and (max-width: 600px) {
  .swiper-image{
    height: 100px;
  }
}
</style> 