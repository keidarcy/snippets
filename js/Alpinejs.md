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

<div x-show.transition="selectedTab === 'tab1'">
  <template :key="index" x-for="(item, index) in items">
    <p>tab1</p>
  </template>
</div>
<div x-show.transition="selectedTab === 'tab2'">
  <template :key="index" x-for="(product, index) in products">
    <p>tab2</p>
  </template>
</div>
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
