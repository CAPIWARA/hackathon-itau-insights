<template>
  <div class="ThemeProgress">
    <p class="value">{{ progress }}</p>
    <figure class="progress">
      <div class="fill" :style="{ width: progress }" />
    </figure>
  </div>
</template>

<script>
  export default {
    props: {
      value: {
        type: Number,
        required: true
      }
    },
    computed: {
      progress () {
        return this.value + '%';
      }
    }
  };
</script>

<style lang="stylus">
  $ThemeProgress-size = 10px;

  .ThemeProgress
    > .value
      margin: 0
      color: #45B69E
      font-size: 12pt
      text-align: center

    > .progress
      height: $ThemeProgress-size
      margin: 0
      border: 1px solid #000
      border-radius: $ThemeProgress-size

      > .fill
        height: 100%
        border-radius: $ThemeProgress-size
        background-image: linear-gradient(90deg, #00C225, #0086C2)
</style>
