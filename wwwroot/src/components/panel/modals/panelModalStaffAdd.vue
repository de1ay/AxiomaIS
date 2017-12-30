<template>
    <form id="panel_modal_franchises">
        <div class="form-field">
            <multiselect
                class="form-field__input"
                v-model="billing_deal"
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
            v-model.trim.number="billing_sum"
            type="number"
            placeholder="Перечислено">
            <icon name="money" scale="1.2" class="form-field__icon--money"></icon>
        </div>
        <div class="form-field">
            <flat-pickr
            v-model="billing_date"
            class="form-field__input"
            :config="date_config"
            placeholder="Дата оплаты"></flat-pickr>
            <icon name="flag-o" scale="1.2" class="form-field__icon--flag"></icon>
        </div>
        <div class="form-field">
            <flat-pickr
            v-model="billing_transfer_date"
            class="form-field__input"
            :config="date_config"
            placeholder="Дата перечисления"></flat-pickr>
            <icon name="flag-checkered" scale="1.2" class="form-field__icon--flag"></icon>
        </div>
        <div class="form-field__actions">
            <input class="form-field__submit" @click.prevent="addBilling" type="submit" value="Создать"/>
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
    name: 'panelModalBillingAdd',
    props: ['franchises', 'clients', 'media', 'deals'],
    data () {
      return {
        billing_deal: null,
        billing_date: '',
        billing_sum: '',
        billing_transfer_date: '',
        date_config: {
          dateFormat: 'd/m/Y',
          locale: Russian
        }
      }
    },
    methods: {
      addBilling () {
        this.$emit('addBilling', {
          billing_deal_info: this.billing_deal,
          billing_date: this.billing_date,
          billing_deal: this.billing_deal.deal_id,
          billing_sum: this.billing_sum,
          billing_transfer_date: this.billing_transfer_date
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

    .form-field__submit {
        width: 100%;
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

    .form-field__submit:hover {  background: #2980b9; }

    .form-field__actions {
        display: flex;
        flex-direction: row;
        justify-content: flex-start;
        align-items: center;
        margin: 15px 0 30px 0;
        width: 300px;
    }

</style>
