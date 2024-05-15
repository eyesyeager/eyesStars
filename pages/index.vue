<template>
  <Sky />
  <DialogTip :show="processStore.isShowTip" :msg="processStore.tipMsg" />
  <DialogNotice :show="processStore.isShowNotice" />
  <DialogPublish :show="processStore.isShowPublish" />
  <DialogStar :show="processStore.isShowStar" />
</template>

<script setup lang="ts">
const contextStore = useContext();
const processStore = useProcess();

useHead({
  title: contextStore.value.appTitle,
  meta: [
    { name: 'keywords', content: contextStore.value.metaKeyWord },
    { name: 'description', content: contextStore.value.metaDesc }
  ],
  link: [{ rel: "icon", type: "image/x-icon", href: contextStore.value.favicon }],
});

</script>

<style>
* {
  margin: 0;
  padding: 0;
}

body {
  background-color: #000;
  overflow: hidden;
}
</style>