<template>
  <div id="panel_main">
    <panel_preloader v-if="!contentLoaded">
    </panel_preloader>
    <panel-modal
      v-if="modal_active"
      :modal_name="modal_name"
      :modal_active.sync="modal_active"
      :franchises.sync="franchises"
      :clinics.sync="clinics"
      :staff.sync="staff"
      :token="token"
      :additional_data="additional_data">
    </panel-modal>
    <div class="menu">
      <div class="brand">
        <span class="brand-text">AxiomaIS</span>
      </div>
      <div class="menu__item"
           :class="{'menu__item--active': selected_module === 'panel_franchises'}"
           @click="selected_module = 'panel_franchises'">
        <icon name="building" scale="1" class="menu__item-icon"></icon>
        <span class="menu__item-text">{{modules_headers_headlines['panel_franchises']}}</span>
      </div>
      <div class="menu__item"
           :class="{'menu__item--active': selected_module === 'panel_clinics'}"
           @click="selected_module = 'panel_clinics'">
        <icon name="plus-circle" scale="1" class="menu__item-icon"></icon>
        <span class="menu__item-text">{{modules_headers_headlines['panel_clinics']}}</span>
      </div>
      <div class="menu__item"
           :class="{'menu__item--active': selected_module === 'panel_staff'}"
           @click="selected_module = 'panel_staff'">
        <icon name="users" scale="1" class="menu__item-icon"></icon>
        <span class="menu__item-text">{{modules_headers_headlines['panel_staff']}}</span>
      </div>
    </div>
    <div class="container">
      <div class="header">
        <div class="header-page_name">
          <span class="page_name-text">{{modules_headers_headlines[selected_module]}}</span>
        </div>
        <div class="header-actions" v-if="selected_module !== 'panel_gantt'">
          <button @click="showModal(selected_module + '_add')" class="actions__button">{{modules_button_text[selected_module]}}</button>
        </div>
      </div>
      <component :is="selected_module"
      :franchises.sync="franchises"
      :clinics.sync="clinics"
      :staff.sync="staff"
      @showModal="showModal"></component>
    </div>
  </div>
</template>

<script>
  import 'vue-awesome/icons/plus-circle'
  import 'vue-awesome/icons/building'
  import 'vue-awesome/icons/users'
  import axios from 'axios'
  import moment from 'moment'
  import panelModal from '@/components/panel/panelModal'
  import panelPreloader from '@/components/preloader'
  import panelFranchises from '@/components/panel/modules/panelFranchises'
  import panelClinics from '@/components/panel/modules/panelClinics'
  import panelStaff from '@/components/panel/modules/panelStaff'
  export default {
    name: 'panelMain',
    data () {
      return {
        token: this.$cookie.get('sid'),
        modal_active: false,
        modal_name: '',
        additional_data: [],
        modules_headers_headlines: {
          'panel_franchises': 'Франчайзи',
          'panel_clinics': 'Клиники',
          'panel_staff': 'Персонал'
        },
        modules_button_text: {
          'panel_franchises': 'Новая франчайзи',
          'panel_clinics': 'Новая клиника',
          'panel_staff': 'Новый сотрудник'
        },
        selected_module: 'panel_franchises',
        franchises: [],
        clinics: [],
        staff: [],
        contentLoaded: false
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
        if (this.$cookie.get('last_page')) {
          this.selected_module = this.$cookie.get('last_page')
        }
        axios.get('https://api.axiomais.ru/franchises?token=' + this.token).then(resp => {
          if (resp.data.code === 200) {
            this.franchises = resp.data.message
          }
        })
        axios.get('https://api.axiomais.ru/clinics?token=' + this.token).then(resp => {
          if (resp.data.code === 200) {
            this.clinics = resp.data.message
          }
          this.contentLoaded = true
        })
      })
    },
    components: {
      panelModal,
      'panel_franchises': panelFranchises,
      'panel_staff': panelStaff,
      'panel_clinics': panelClinics,
      'panel_preloader': panelPreloader
    },
    methods: {
      showModal (modalName, additionalData) {
        this.modal_name = modalName
        this.modal_active = true
        this.additional_data = additionalData
      },
      parseMedia () {
        this.media.forEach(item => {
          item.media_type_info = {
            value: item.media_type,
            label: this.parseMediaType(item.media_type)
          }
        })
      },
      parseMediaType (mediaType) {
        switch (mediaType) {
          case '1': return 'Баннер'
          case '2': return 'Радиостанция'
          default: return 'Ошибка'
        }
      },
      parseDealType (dealType) {
        switch (dealType) {
          case '0': return {
            type_id: dealType,
            type_name: 'Размещение'
          }
          case '1': return {
            type_id: dealType,
            type_name: 'Сбыт'
          }
          case '2': return {
            type_id: dealType,
            type_name: 'Изготовление'
          }
          case '3': return {
            type_id: dealType,
            type_name: 'Реализация'
          }
          case '4': return {
            type_id: dealType,
            type_name: 'Бартер'
          }
          default: return 'Ошибка!'
        }
      },
      parseBillings () {
        this.franchises.forEach(billing => {
          billing.billing_deal_info = this.deals.filter(deal => {
            return deal.deal_id === billing.billing_deal
          })[0]
        })
      },
      parseDeals () {
        this.deals.forEach(deal => {
          deal.deal_client_info = this.findClientByID(deal.deal_client)
          deal.deal_media_info = this.findMediaByID(deal.deal_media)
          deal.deal_type_info = this.parseDealType(deal.deal_type)
          let dealPeriods = deal.deal_period.split(';')
          deal.deal_periods = []
          let dates = []
          dealPeriods.forEach((dealPeriod, index) => {
            if (dealPeriod.length > 1) {
              dealPeriod = dealPeriod.split('-')
              dates.push(dealPeriod[0])
              dates.push(dealPeriod[1])
              deal.deal_periods.push({
                period_start: dealPeriod[0],
                period_end: dealPeriod[1],
                period_date: dealPeriod[0] + '-' + dealPeriod[1]
              })
            }
          })
          dates = dates.sort((date1, date2) => {
            return moment(date1, 'DD/MM/YYYY').toDate() > moment(date2, 'DD/MM/YYYY').toDate()
          })
          deal.start_date = dates[0]
          deal.end_date = dates[dates.length - 1]
          if (deal.deal_status === '2' && new Date() > moment(deal.start_date, 'DD/MM/YYYY').toDate() && new Date() < moment(deal.end_date, 'DD/MM/YYYY').toDate()) {
            deal.deal_status = '3'
            axios.put('https://beta.project.nullteam.info/api/deals/' + deal.deal_id, deal)
          }
          if (deal.deal_status !== '4' && new Date() > moment(deal.end_date, 'DD/MM/YYYY').toDate()) {
            deal.deal_status = '4'
            axios.put('https://beta.project.nullteam.info/api/deals/' + deal.deal_id, deal)
          }
        })
      },
      findClientByID (clientId) {
        try {
          return this.clients.filter((client, index) => {
            return client.client_id === clientId
          })[0]
        } catch (e) {
          return {
            client_id: 0,
            client_name: 'Клиент не найден'
          }
        }
      },
      findMediaByID (mediaId) {
        try {
          return this.media.filter((media, index) => {
            return media.media_id === mediaId
          })[0]
        } catch (e) {
          return {
            media_id: 0,
            media_name: 'Медиа-носитель не найден',
            media_type: 0,
            media_address: ''
          }
        }
      }
    },
    watch: {
      selected_module: function (newSelect) {
        this.$cookie.set('last_page', newSelect, 365)
      }
    }
  }
</script>

<style scoped="true">

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
