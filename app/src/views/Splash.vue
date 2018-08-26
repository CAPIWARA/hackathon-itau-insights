<template>
  <section>
    <h2>Bem-vindo ao</h2>
    <h1>Nome do Aplicativo</h1>
    <theme-button link="/profile">Começar</theme-button>
  </section>
</template>

<script>
  import ThemeButton from '@/components/Theme/ThemeButton';

  export default {
    components: { ThemeButton }
  };
</script>
