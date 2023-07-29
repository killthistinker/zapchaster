<template>
    <div class="product">
    <v-popup
        v-if="isProductInfoVisible"
        @closePopup="closePopup"
        :partDetail="partDetail"
    />
     <div class="image"><img alt="photo" :src="'data:image/png;base64,'+ product_data.mainPhoto "></div>
     <div class="product-footer">
        <div v-if="product_data.price !== 0" class="price"><p class="text">Цена: {{product_data.price}}</p></div>
        <div class="description"><p class="text">Винкод: {{product_data.winCode}}</p></div>
        <div class="description"><p class="text">{{product_data.title}}</p></div>
         <div class="button-block">
          <a @click="counter" :href="`https://wa.me/77078276330?text=Здравствуйте,%20по%20поводу%20детали%20${product_data.title}%20код%20детали:%20${product_data.winCode}`" class="buy-button">Заказать</a>
            <button 
            class="product-button"
            @click="showProductInfo"
            >
            Информация
            </button>
        </div>
     </div>
    </div>
</template>
   
<script>
   import vPopup from '../components/popups/v-popup.vue'
   import { getProductData } from '@/api/parts'
   import { counter } from '@/api/parts'

   export default {
       name: 'Product-vue',
       components:{
       vPopup,
       },
       props: {
        product_data:{
            type: Object,
            default() {
                return {}
            }
        }
       },
       data() { 
        return {
            partDetail: {
             
            },
            isProductInfoVisible: false,
        } 
     },
      methods: {
       async showProductInfo(){
           const response = await getProductData(this.product_data.id)
           console.log(response.data)
           this.partDetail = response.data
           this.isProductInfoVisible = true
        },
        closePopup(){
            this.isProductInfoVisible = false
        },

        async counter(){
          var data = {
            'winCode': this.product_data.winCode,
            'count': 0
          }
           const response = await counter(data)
           console.log(response)
        }
 }
}

</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@300;400;500;600&display=swap');
  .product{
    width: 32em;
    display: flex;
    color: aliceblue;
    margin: 20px 0px;
    border: 1px solid gray;
    border-radius: 5px;
   }

   .image{
    width: 25.5em;
    height: 100%;
   }

   .product-button{
    margin: 0 5px;
    width: 8em;
    height: 2em;
    border: 2px solid #736161;
    border-radius: 8px 8px;
    background-color: #d2d2d2;
    cursor: pointer;
    transition: all 0.3s ease
   }

   .product-button:hover{
    filter: saturate(250%);
   }

   .product-image{
      width: 100%;
   }

   .product-footer{
    width: 100%;
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    background-color: #69696991;
    color: #ffffff7d;
    font-size: 20px;
    font-family: 'Montserrat';
    font-weight: 500;
    text-align: left;
    padding-left: 15px;
   }

   .button-block{
    display: flex;
    align-items: center;
    justify-content: center;
   }

   .buy-button{
    padding: 0px 10px;
    text-decoration: none;
    color: black;
    background-color: #fff1de;
    transition: all 0.3s ease;
    font-family: 'Montserrat';
    border: 2px solid #736161;
    border-radius: 8px 8px;
   }

   .buy-button:hover{
    background-color: #a78383;
   }

   
@media screen and (max-width: 1280px) {
  .product {
    width: 16em;
  }

  .image{
    width: 16.5em;
  }

  .product-footer{
    font-size: 15px;;
  }
}

@media screen and (max-width: 600px) {
  .product {
    width: 100%;
  }
}

</style>
   