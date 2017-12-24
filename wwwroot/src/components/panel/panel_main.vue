<template>
  <div id="panel_main">
    <panel-preloader v-if="!contentLoaded"></panel-preloader>
    <div class="menu">
      <div class="brand">
        <span class="brand-text">AxiomaIS</span>
      </div>
      <div class="menu__item"
           :class="{'menu__item--active': selected_module === 'panel_roles'}"
           @click="selected_module = 'panel_roles'">
        <icon name="user" scale="1" class="menu__item-icon"></icon>
        <span class="menu__item-text">{{modules_headers_headlines['panel_roles']}}</span>
      </div>
    </div>
    <div class="container">
      <div class="header">
        <div class="header-page_name">
          <span class="page_name-text">{{modules_headers_headlines[selected_module]}}</span>
        </div>
        <div class="header-actions">
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import axios from 'axios'
  import PanelPreloader from '../preloader'

  export default {
    components: {PanelPreloader},
    name: 'PanelMain',
    data () {
      return {
        contentLoaded: false,
        token: this.$cookie.get('sid'),
        selected_module: 'panel_roles',
        modules_headers_headlines: {
          'panel_roles': 'Роли'
        }
      }
    },
    created () {
      axios.get('https://api.axiomais.ru/token.verify?token=' + this.token).then(resp => {
        if (resp.data.code !== 200) {
          this.$cookie.set('sid', '', -1)
          this.$snotify.error(resp.data.message, 'Ошибка!', {
            timeout: 2000,
            showProgressBar: true,
            closeOnClick: true,
            pauseOnHover: true
          })
          this.$router.push({name: 'sign_in'})
        }
        this.contentLoaded = true
      })
    }
  }
</script>

<style scoped>

  #panel_main {
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: stretch;
    width: 100%;
    height: 100vh;
    max-height: 100vh;
    overflow: hidden;
  }

  .menu {
    z-index: 999;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: center;
    height: 100%;
    min-width: 200px;
    background: #fff;
    box-shadow: 0.5px 0 1px 0 rgba(50, 50, 50, 0.3);
  }

  .container {
    z-index: 10;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: stretch;
    flex-grow: 1;
    overflow-y: scroll;
  }

  .header {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    height: 50px;
    width: 100%;
    background: #fff;
    box-shadow: 0 0.5px 1px 0 rgba(50, 50, 50, 0.3);
  }

  .header-page_name { margin-left: 50px; }

  .page_name-text {
    font-size: 24px;
    color: #3498db;
  }

  .header-actions {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    margin-right: 50px;
  }

  .actions__button {
    padding: 5px 10px;
    color: #fff;
    background: #3498db;
    font-size: 18px;
    border-radius: 10px;
    border: none;
    outline: none;
    transition: all 0.3s ease-in-out;
  }

  .actions__button:hover {
    cursor: pointer;
    background: #2980b9;
  }

  .brand {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    height: 50px;
    width: 100%;
    box-shadow: 0 0.5px 1px 0 rgba(50, 50, 50, 0.3);
  }

  .brand-text {
    font-size: 28px;
    font-weight: 800;
    color: #3498db;
  }

  .menu__item:nth-child(2) { margin-top: 50px; }

  .menu__item {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    height: 50px;
    width: 90%;
    margin-left: 10%;
    color: #3498db;
    font-size: 24px;
    transition: all 0.3s ease-in-out;
  }

  .menu__item:hover { cursor: pointer; }

  .menu__item:not(.menu__item--active):hover {
    color: #fff;
    background: #3498db;
  }

  .menu__item--active {
    color: #fff;
    background: #3498db;
  }

  .menu__item--active:after {
    content: '';
    position: absolute;
    left: 190px;
    width: 10px;
    height: 50px;
    background: #2980b9;
  }

  .menu__item-icon {
    margin-top: 4px;
    margin-right: 10px;
  }

  .menu__item-text { text-transform: capitalize; }

</style>
