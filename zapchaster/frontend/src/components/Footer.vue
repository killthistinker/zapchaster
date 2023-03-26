<template>
    <footer class="site-footer">
      <div class="container">
        <div class="row">
          <div class="col-md-4">
            <h3>Контакты</h3>
            <p>Адрес: г. Алматы</p>
            <p>Телефон: +7 (999) 123-45-67</p>
            <p>Email: info@autostore.com</p>
          </div>
        </div>
        <hr>
        <div class="row">
          <div class="col-md-12">
            <p class="text-center">&copy; {{ currentYear }} Autostore. Все права защищены.</p>
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
 }

 .links{
    text-decoration: none;
    color: aliceblue;
 }
</style>