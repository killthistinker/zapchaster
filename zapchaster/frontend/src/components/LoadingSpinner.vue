<!-- eslint-disable -->
<template>
    <div class="preLoader">
      <!-- Анимация загрузки -->
      <div class="spinner"></div>
    </div>
  </template>
  
  <style>
  .preLoader {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
  }
  
  .spinner {
    width: 100px;
    height: 100px;
    border: 4px solid #ffffff;
    border-top: 4px solid #3a3d40;
    border-radius: 50%;
    animation: spin 1s linear infinite;
  }
  
  @keyframes spin {
    0% { transform: rotate(0deg); }
    100% { transform: rotate(360deg); }
  }
  </style>
  