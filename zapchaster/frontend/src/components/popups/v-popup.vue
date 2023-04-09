<template>
    <div class="v-popup">
        <div class="v-popup__header">
            <span>
                <i @click="closePopup" class="material-icons">close</i>
            </span>
        </div>
        <div class="v-popup__content">
            <slot>
                <div class="content__image">
                    <img src="../../../public/background.jpg" alt="">
                </div>
                <div class="title"><h1>Lorem ipsum dolor sit amet.</h1></div>
                <div class="discription">
                    <p>Lorem ipsum dolor, sit amet consectetur adipisicing elit. Molestiae quia doloribus illum, voluptates consequuntur praesentium natus repellat suscipit consequatur, voluptas necessitatibus at eum officiis fuga provident magni quibusdam cumque maiores!</p>
                </div>
            </slot>
        </div>
        <div class="v-popup__footer">
            <button class="add__product btn">Купить</button>
            <button @click="closePopup" class="close__popup">Закрыть</button>
        </div>
    </div>
</template>

<script>
    export default{
        name: 'v-popup',
        methods: {
            closePopup (){
                this.$emit('closePopup')
            }
        },
    }

</script>

<style>
    .v-popup{
        padding: 20px 22px;
        height: 80%;
        width: 35%;
        position: fixed;
        left: 50%;
        top: 50%;
        transform:translate(-50%,-50%);
        background-color: #949c9c;
        z-index: 9999;
        display: flex;
        flex-direction: column;
        align-items: center;
        flex-wrap: wrap;
        justify-content: space-between;
        border: 3px solid #ededed;
        border-radius: 5px;
    }

    .v-popup__content{
        height: 80%;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: space-around;
    }

    .content__image{
        width: 300px;
        height: 200px;
    }
</style>