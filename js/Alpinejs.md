## Switch tabs

```html
<ul class="category">
  <li :class="{ 'on': selectedTab === 'tab1' }" @click="handleTabClicked('tab1')">
    <a href="#">tab1</a>
  </li>
  <li :class="{ 'on': selectedTab === 'tab2' }" @click="handleTabClicked('tab2')">
    <a href="#">tab2</a>
  </li>
</ul>

<template x-if.transition.in="selectedTab === 'tab1'">
  <p>tab1</p>
</template>
<template x-if.transition.in="selectedTab === 'tab2'">
  <p>tab2</p>
</template>
<script>
  function data() {
    return {
      selectedTab: 'tab1',
      handleTabClicked(tab) {
        this.selectedTab = tab;
      }
    };
  }
</script>
```
