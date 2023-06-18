<template>
 <Layout/>
    <div class="about-info">
        <h1 class="about-title">Немного о нас</h1>
        <p class="about-description">
            Хотите обновить и улучшить работу своего автомобиля Kia или Hyundai? Мы предлагаем вам широкий выбор качественных запчастей, которые помогут вам достичь оптимальной производительности. Добро пожаловать в нашу компанию!
            Мы специализируемся на продаже запчастей для двигателей автомобилей Kia и Hyundai. У нас вы найдете все необходимое для того, чтобы ваш двигатель работал на высшем уровне. Мы предлагаем поршни, кольца и вкладыши высокого качества, которые гарантируют надежность и долговечность.
        </p>
        <p class="about-description">
            Наши запчасти производятся с использованием передовых технологий и материалов, чтобы обеспечить оптимальную производительность и эффективность вашего двигателя. Мы тщательно подбираем наши запчасти, чтобы удовлетворить самые высокие стандарты качества.
            Наша команда профессионалов всегда готова помочь вам с выбором правильных запчастей и ответить на все ваши вопросы. Мы предлагаем удобный онлайн-каталог, где вы можете легко найти нужные запчасти, сравнить цены и оформить заказ. Мы также гарантируем быструю доставку, чтобы вы могли получить свои запчасти в кратчайшие сроки.
        </p>
        <p class="about-description">
            Не откладывайте на потом замену изношенных запчастей. Обратитесь в нашу компанию и доверьтесь нам, чтобы обеспечить вашему автомобилю надежность, производительность и долговечность. Выберите качество и надежность - выберите нас, вашего надежного партнера по запчастям для двигателей автомобилей Kia и Hyundai.
        </p>    
    </div>
</template>

<script>
import Layout from '../Layouts/Layout.vue'
export default {
    name: 'CompanyInfo',
    components: {
        Layout,
    },
};
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Russo+One&family=Sofia+Sans+Extra+Condensed:wght@500&display=swap');
    .about-info{
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        height: 500px;
        margin: 0 auto;
        margin-top: 5%;
        width: 85%;
        height: 500px;
    }


    .about-title{
        color: #d4d4d4;
        margin: 0 auto;
        font-family: 'Russo One', sans-serif;
        font-weight: 400;
    }

    .about-info{
        margin: 0 auto;
        margin-top: 5%;
        width: 85%;
        background-color: #2e2d2d9e;
        border: 0px solid;
        border-radius: 23px;
    }

    .about-description{
        color: #cfcfcf;
        line-height: 30px;
        font-size: 20px;
        margin: 0 auto;
        font-family: 'Sofia Sans Extra Condensed', sans-serif;
        font-size: 35px;
    }
</style>
