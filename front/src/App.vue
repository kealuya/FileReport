<template>

  <BaseHeader/>
  <div class="main-window">
    <div style="width: 80%">
      <router-view/>
    </div>
  </div>

</template>

<style>
.main-window {
  width: 100vw;
  display: flex;
  flex-direction: column;
  align-items: center;
}
</style>
