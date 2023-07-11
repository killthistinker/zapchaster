<template>
    <footer class="site-footer">
      <div class="container">
        <div class="row">
          <div class="col-md-4">
            <h3>Контакты</h3>
            <p>Адрес: Казахстан, г. Алматы</p>

            <a class="ws__link" href="https://wa.me/77479656291?text=Здраствуйте,%20по%20поводу%20объявления%20с%20Autostore ">Whatsapp: +7 707 827 6330</a>
            <p>Email: popovich74@mail.ru</p>
          </div>
        </div>
        <hr>
        <div class="row"> 
          <div class="col-md-12">
            <p class="text-center">&copy; {{ currentYear }} Zapchastik.kz. Все права защищены.</p>
          </div>
        </div>
      </div>
    </footer>
  </template>
  
  <script>
  export default {
    name: 'FooTer',
    data() {
      return {
        currentYear: new Date().getFullYear()
      }
    }
  }
  </script>

 <style>
 .row{
    display: flex;
    justify-content: space-between;
    color: aliceblue;
    font-family: 'Montserrat', sans-serif;
 }

 .links{
    text-decoration: none;
    color: aliceblue;
 }

 .ws__link{
  text-decoration: none;
  color: rgb(49 247 145 / 84%);
 }
</style>