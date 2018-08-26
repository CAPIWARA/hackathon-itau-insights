<template>
  <section class="ThemeLayout">
    <main class="main">
      <slot />
    </main>

    <nav class="menu">
      <router-link
        to="/supporter"
        class="link"
        active-class="-active"
      >
        <img src="@/assets/icons/Donate@2x.png" class="icon" />
        <span class="text">Apoiar</span>
      </router-link>

      <router-link
        to="/profile"
        class="link"
        active-class="-active"
      >
        <img src="@/assets/icons/House@2x.png" class="icon" />
        <span class="text">Início</span>
      </router-link>

      <router-link
        to="/entrepreneur"
        class="link"
        active-class="-active"
      >
        <img src="@/assets/icons/PigBank@2x.png" class="icon" />
        <span class="text">Campanha</span>
      </router-link>
    </nav>
  </section>
</template>

<style lang="stylus">
  $ThemeLayout-menu-size = 50px
  $ThemeLayout-menu-border-color = rgba(#646464, .6)

  .ThemeLayout
    > .main
      margin-bottom: $ThemeLayout-menu-size

    > .menu
      // Posição
      position: fixed
      left: 0
      right: 0
      bottom: 0

      // Box
      width: 100%
      height: $ThemeLayout-menu-size
      display: flex
      justify-content: space-around
      background-color: #FFF

      > .link
        width: calc(100% / 3)
        height: $ThemeLayout-menu-size
        display: flex
        align-items: center
        flex-direction: column
        justify-content: space-around
        border-left: .5px solid $ThemeLayout-menu-border-color
        border-right: @border-left

        // Aparência
        cursor: pointer
        outline: none
        opacity: .6
        transition: opacity .3s ease-in
        text-decoration: none

        &.-active
          opacity: 1

        > .icon
          align-self: center
          height: 20px

        > .text
          color: #000
</style>
