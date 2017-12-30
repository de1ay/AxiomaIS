<template>
    <form id="panel_modal_franchises">
        <div class="form-field">
            <multiselect
                class="form-field__input"
                v-model="editable_data.billing_deal_info"
                :options="deals"
                label="deal_brand"
                track-by="deal_id"
                :allow-empty="false"
                :show-labels="false"
                placeholder="Сделка"></multiselect>
            <icon name="handshake-o" scale="1.2" class="form-field__icon--handshake"></icon>
        </div>
        <div class="form-field">
            <input
            class="form-field__input"
            v-model.trim.number="editable_data.billing_sum"
            type="number"
            placeholder="Перечислено">
            <icon name="money" scale="1.2" class="form-field__icon--money"></icon>
        </div>
        <div class="form-field">
            <flat-pickr
            v-model="editable_data.billing_date"
            class="form-field__input"
            :config="date_config"
            placeholder="Дата оплаты"></flat-pickr>
            <icon name="flag-o" scale="1.2" class="form-field__icon--flag"></icon>
        </div>
        <div class="form-field">
            <flat-pickr
            v-model="editable_data.billing_transfer_date"
            class="form-field__input"
            :config="date_config"
            placeholder="Дата перечисления"></flat-pickr>
            <icon name="flag-checkered" scale="1.2" class="form-field__icon--flag"></icon>
        </div>
        <div class="form-field__actions">
          <button class="form-field__edit" @click.prevent="editBillingConfirm">Изменить</button>
          <button class="form-field__delete" @click.prevent="deleteBillingConfirm">Удалить</button>
        </div>
    </form>
</template>

<script>
  import 'vue-awesome/icons/handshake-o'
  import 'vue-awesome/icons/money'
  import 'vue-awesome/icons/flag-o'
  import 'vue-awesome/icons/flag-checkered'
  import flatPickr from 'vue-flatpickr-component'
  import 'flatpickr/dist/flatpickr.css'
  import {Russian} from 'flatpickr/dist/l10n/ru'
  export default {
    name: 'panelModalBillingView',
    props: ['franchises', 'clients', 'media', 'deals', 'additional_data'],
    data () {
      return {
        editable_data: Object.assign({}, this.additional_data),
        date_config: {
          dateFormat: 'd/m/Y',
          locale: Russian
        }
      }
    },
    methods: {
      editBillingConfirm () {
        this.$snotify.confirm('Изменить оплату?', 'Изменение', {
          timeout: 2000,
          showProgressBar: true,
          closeOnClick: false,
          pauseOnHover: true,
          buttons: [
            {text: 'Да', action: (notifyId) => { this.editBilling(); this.$snotify.remove(notifyId) }},
            {text: 'Нет', action: (notifyId) => this.$snotify.remove(notifyId)}
          ]
        })
      },
      editBilling () {
        this.editable_data.billing_deal = this.editable_data.billing_deal_info.deal_id
        this.$emit('editBilling', this.editable_data, this.additional_data)
      },
      deleteBillingConfirm () {
        this.$snotify.confirm('Удалить оплату?', 'Удаление', {
          timeout: 2000,
          showProgressBar: true,
          closeOnClick: false,
          pauseOnHover: true,
          buttons: [
            {text: 'Да', action: (notifyId) => { this.$emit('deleteBilling', this.additional_data); this.$snotify.remove(notifyId) }},
            {text: 'Нет', action: (notifyId) => this.$snotify.remove(notifyId)}
          ]
        })
      }
    },
    components: {
      flatPickr
    }
  }
</script>

<style src="vue-multiselect/dist/vue-multiselect.min.css"></style>

<style scoped>

    #panel_modal_franchises {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    }

    .form-field__delete {
        width: 45%;
        height: 38px;
        border-radius: 5px;
        border: none;
        background-color: #e74c3c;
        text-transform: uppercase;
        color: #fff;
        font-size: 18px;
        outline: none;
        cursor: pointer;
        transition: all 0.3s ease-in-out;
    }

    .form-field__delete:hover {  background: #c0392b; }

    .form-field__edit {
        width: 45%;
        height: 38px;
        border-radius: 5px;
        border: none;
        background-color: #3498db;
        text-transform: uppercase;
        color: #fff;
        font-size: 18px;
        outline: none;
        cursor: pointer;
        transition: all 0.3s ease-in-out;
    }

    .form-field__edit:hover {  background: #2980b9; }

    .form-field__actions {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        margin: 15px 0 30px 0;
        width: 300px;
    }

</style>
