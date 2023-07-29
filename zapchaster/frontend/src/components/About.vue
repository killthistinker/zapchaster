<template>
 <Layout/>
    <div class="about-info">
        <h1 class="about-title">Добро пожаловать в нашу компанию!</h1>
        <p class="about-description-left-block">
            <b>Мы предлагаем</b> поршни, кольца и вкладыши высокого качества, которые гарантируют надежность и долговечность.
        </p>
        <p class="about-descriptio-bottom">
            Наши запчасти производятся с использованием передовых технологий и материалов, чтобы обеспечить оптимальную производительность и эффективность вашего двигателя. Мы тщательно подбираем наши запчасти, чтобы удовлетворить самые высокие стандарты качества.
            Наша команда профессионалов всегда готова помочь вам с выбором правильных запчастей и ответить на все ваши вопросы. Мы предлагаем удобный онлайн-каталог, где вы можете легко найти нужные запчасти, сравнить цены и оформить заказ. Мы также гарантируем быструю доставку, чтобы вы могли получить свои запчасти в кратчайшие сроки.
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
@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@300;400;500;600&display=swap');
    .about-info{
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        margin: 0 auto;
        margin-top: 5%;
        width: 85%;
        height: 100%;
    }


    .about-title{
        display: flex;
        color: #d4d4d4;
        margin: 0 2em;
        font-family: 'Montserrat', sans-serif;
        font-weight: 400;
        flex-direction: row;
        justify-content: flex-end;
        text-align: right;
        font-size: 72px;
    }

    .about-description-left-block{
        color: #d9aa83c7;
        line-height: 40px;
        margin: 30px 0;
        font-family: 'Montserrat', sans-serif;
        font-size: 52px;
        text-align: left;
        width: 76%;
    }

    .about-descriptio-bottom{
        font-size: 17px;
        line-height: 28px;
        color: #d7d7d7;
        font-family: 'Montserrat', sans-serif;
    }

    @media screen and (max-width: 600px) {
    .about-title{
        font-size: 32px;
        margin: 0 0;
    }
    .about-description-left-block{
        width: 100%;
        line-height: 25px;
        font-size: 26px;
    }

    .about-descriptio-bottom{
        font-size: 14px;
        line-height: 22px;
    }
}
</style>
