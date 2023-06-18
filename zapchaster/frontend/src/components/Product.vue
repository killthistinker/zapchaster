<template>
    <div class="product">
    <v-popup
        v-if="isProductInfoVisible"
        @closePopup="closePopup"
        :partDetail="partDetail"
    />
     <div class="image"><img alt="photo" :src="'data:image/png;base64,'+ product_data.mainPhoto "></div>
     <div class="product-footer">
        <div class="price"><p class="text">Цена: {{product_data.price}}</p></div>
        <div class="description"><p class="text">Винкод: {{product_data.winCode}}</p></div>
        <div class="description"><p class="text">{{product_data.title}}</p></div>
         <div class="button-block">
            <a :href="'https://wa.me/77479656291?text=Здравствуйте,%20по%20поводу%20детали%20код%20детали:%20' + product_data.winCode" class="buy-button">Заказать</a>
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
        showProductInfo(){
           this.$axios.get(`http://localhost:8080/detail?partId=${this.product_data.id}`).then(response => (this.partDetail = response.data))
           this.isProductInfoVisible = true
        },
        closePopup(){
            this.isProductInfoVisible = false
        }
 }
}

</script>

<style>
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
    background-color: #ffecd2;
    cursor: pointer;
    transition: all 0.3s ease;
    font-family: 'Palanquin Dark', sans-serif;
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
    background-color: #a89f9fbf;
    color: black;
    font-size: 17px;
    font-family: 'Unbounded';
   }

   .button-block{
    display: flex;
    align-items: center;
    justify-content: center;
   }

   .buy-button{
    padding: 1px 14px;
    text-decoration: none;
    color: black;
    background-color: #ffecd2;
    transition: all 0.3s ease;
    font-family: 'Palanquin Dark', sans-serif;
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
   