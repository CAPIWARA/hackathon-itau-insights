<template>
  <component
    class="ThemeButton"
    :is="isLink ? 'router-link' : 'button'"
    :to="isLink ? link : undefined"
    @click="$emit('click', $event)"
  >
    <span class="text">
      <slot />
    </span>
  </component>
</template>

<script>
  import { isNotEmptyString } from '@/domains/validators';

  export default {
    props: {
      link: {
        type: String,
        default: null
      }
    },
    computed: {
      isLink () {
        return isNotEmptyString(this.link);
      }
    }
  };
</script>
