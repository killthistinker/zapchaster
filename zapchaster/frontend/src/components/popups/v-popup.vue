<template>
    <div class="v-popup">
       <div class="v-popup-content">
        <div class="v-popup__header">
            <span>
                <svg class="close-btn" @click="closePopup" xmlns="http://www.w3.org/2000/svg" height="48" viewBox="0 96 960 960" width="48"><path d="m249 849-42-42 231-231-231-231 42-42 231 231 231-231 42 42-231 231 231 231-42 42-231-231-231 231Z"/></svg>
            </span>
        </div>
        <div class="v-popup__content">
            <slot>
                <div class="image-block">
                    <div class="content__image">
                <VpopupSlider :partImages="partDetail.partImages"/>
                </div>
                </div>
                <div class="main-info">
                    <div class="title"><h1>{{partDetail.title}}</h1></div>
                <div class="price"><p>Цена: {{partDetail.price}}</p></div>
                <div class="manufacter"><p>Производитель: {{partDetail.manufacturer}}</p></div>
                <div class="brand"><h1>Марка: {{partDetail.carBrand}}</h1></div>
                </div>
                <div class="discription">
                    <p>{{partDetail.description}}</p>
                </div>
            </slot>
        </div>
        <div class="v-popup__footer">
            <a @click="counter" :href="`https://wa.me/77078276330?text=Здравствуйте,%20по%20поводу%20детали%20${partDetail.title}%20код%20детали:%20${partDetail.winCode}`" class="buy-button">Заказать</a>
            <button @click="closePopup" class="close__popup btn">Закрыть</button>
            
        </div>
       </div>
    </div>
</template>

<script>
import VpopupSlider from '../sliders/v-popup-slider.vue'
import { counter } from '@/api/parts'
    export default{
        name: 'v-popup',
        components:{
            VpopupSlider
        },
        props: {
            partDetail:{
            type: Object,
            default() {
                return {}
            },
        },
       },
        methods: {
            closePopup (){
                this.$emit('closePopup')
            },

        async counter(){
          var data = {
            'winCode': this.partDetail.winCode,
            'count': 0
          }
           const response = await counter(data)
           console.log(response)
        }

        },  
    };
</script>

<style>
    .v-popup{
    height: 80%;
    padding: 20px 22px;
    width: 35%;
    position: fixed;
    left: 50%;
    top: 50%;
    transform: translate(-50%,-50%);
    background-color: #69696991;
    z-index: 9999;
    border: 1px solid #acaaaa;
    border-radius: 0.5em;
    color: black;
    overflow: auto;
    }

    .v-popup-content{
    height: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    flex-wrap: wrap;
    justify-content: space-evenly;
    font-size: 15px;
    font-family: 'Unbounded';
    }

    .v-popup__header{
     position: fixed;
     top: 0.5em;
     right: 1.5em;
    }

    .v-popup__content{
        height: 90%;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: space-around;
    }

    .content__image{
        margin: 0 auto;
    width: 400px;
    height: 200px;
    margin-bottom: 1em;
    }

    .image-block{
        width: 100%;
    }

    .v-popup__footer{
    display: flex;
    width: 100%;
    justify-content: center;
    }

    .title{
        padding: 0.5em 0em;
    }

   .btn{
    margin: 0 5px;
    width: 8em;
    height: 2em;
    border: 2px solid #736161;
    border-radius: 8px 8px;
    cursor: pointer;
    background-color: #F99F27;
    transition: all 0.3s ease;
    cursor: pointer;
    font-family: 'Palanquin Dark', sans-serif;
     }

   .btn:hover{
    filter: saturate(250%);
   }

   .close-btn{
    cursor: pointer;
   }

   .main-info{
    width: 85%;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    font-size: 12px;
   }

   .price{
    line-height: 5px;
   }

   .manufacter{
    line-height: 5px;
   }

   @media screen and (max-width: 600px) {
    .modal {
      width: 90%;
      height: 90%;
    }

    .v-popup{
    width: 80%;
    }

    .v-popup{
        font-size: 12px;
    }

    .v-popup-content {
    height: 100%;
    }
  
    .content__image {
    width: 325px;
    height: 200px;
}
  }
</style>