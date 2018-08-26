<template>
  <section class="ThemeLayout">
    <main class="main">
      <slot />
    </main>
    <nav class="menu">
      <router-link to="/supporter">Apoiador</router-link>
      <router-link to="/profile">Perfil</router-link>
      <router-link to="/entrepreneur">Empreendedor</router-link>
    </nav>
  </section>
</template>
