<template>
  <form id="panel_modal_franchises">
    <div class="form-field">
        <input
        class="form-field__input"
        v-model.trim="name"
        type="text"
        placeholder="Название">
        <icon name="building" scale="1.2"></icon>
    </div>
    <div class="form-field">
      <input
        class="form-field__input"
        v-model.trim="official_name"
        type="text"
        placeholder="Официальное наименование">
      <icon name="building" scale="1.2"></icon>
    </div>
    <div class="form-field">
      <input
        class="form-field__input"
        v-model.trim="full_official_name"
        type="text"
        placeholder="Полное официальное наименование">
      <icon name="building" scale="1.2"></icon>
    </div>
    <div class="form-field">
      <input
        class="form-field__input"
        v-model.trim="license"
        type="text"
        placeholder="Лицензия">
      <icon name="id-card" scale="1.2"></icon>
    </div>
    <div class="form-field">
      <input
        class="form-field__input"
        v-model.trim="city"
        type="text"
        placeholder="Город">
      <icon name="industry" scale="1.2"></icon>
    </div>
    <div class="form-field">
      <input
        class="form-field__input"
        v-model.trim="legal_address"
        type="text"
        placeholder="Юридический адрес">
      <icon name="legal" scale="1.2"></icon>
    </div>
    <div class="form-field">
      <input
        class="form-field__input"
        v-model.trim="inn"
        type="text"
        pattern="\[0-9]{10}"
        maxlength="10"
        placeholder="ИНН">
      <icon name="ticket" scale="1.2"></icon>
    </div>
    <div class="form-field">
      <input
        class="form-field__input"
        v-model.trim="kpp"
        type="text"
        pattern="\[0-9]{9}"
        maxlength="9"
        placeholder="КПП">
      <icon name="sticky-note" scale="1.2"></icon>
    </div>
    <div class="form-field">
      <input
        class="form-field__input"
        v-model.trim="bank_account"
        type="text"
        pattern="\[0-9]{20}"
        maxlength="20"
        placeholder="Р/C">
      <icon name="credit-card" scale="1.2"></icon>
    </div>
    <div class="form-field">
      <input
        class="form-field__input"
        v-model.trim="bank"
        type="text"
        placeholder="Банк">
      <icon name="bank" scale="1.2"></icon>
    </div>
    <div class="form-field">
      <input
        class="form-field__input"
        v-model.trim="correspondent_account"
        type="text"
        pattern="\[0-9]{19}"
        maxlength="19"
        placeholder="Кор. счёт">
      <icon name="money" scale="1.2"></icon>
    </div>
    <div class="form-field">
      <input
        class="form-field__input"
        v-model.trim="bik"
        type="text"
        pattern="\[0-9]{9}"
        maxlength="9"
        placeholder="БИК">
      <icon name="exchange" scale="1.2"></icon>
    </div>
    <div class="form-field">
      <input
        class="form-field__input"
        v-model.trim="ogrn"
        type="text"
        pattern="\[0-9]{13}"
        maxlength="13"
        placeholder="ОГРН">
      <icon name="tag" scale="1.2"></icon>
    </div>
    <div class="form-field">
      <input
        class="form-field__input"
        v-model.trim="okved"
        type="text"
        placeholder="ОКВЭД">
      <icon name="info-circle" scale="1.2"></icon>
    </div>
    <div class="form-field">
      <input
        class="form-field__input"
        v-model.trim="act"
        type="text"
        placeholder="Действует">
      <icon name="registered" scale="1.2"></icon>
    </div>
    <div class="form-field">
      <input
        class="form-field__input"
        v-model.trim="correspondent_address"
        type="text"
        placeholder="Адрес для корреспонденции">
      <icon name="map" scale="1.2"></icon>
    </div>
    <div class="form-field">
      <the-mask
        class="form-field__input"
        v-model.trim="number"
        mask="+7 (###) ###-##-##"
        type="tel"
        placeholder="Телефон"></the-mask>
      <icon name="phone" scale="1.2"></icon>
    </div>
    <div class="form-field__actions">
      <input class="form-field__submit" @click.prevent="addFranchise" type="submit" value="Создать"/>
    </div>
  </form>
</template>

<script>
  import 'vue-awesome/icons/building'
  import 'vue-awesome/icons/id-card'
  import 'vue-awesome/icons/industry'
  import 'vue-awesome/icons/map'
  import 'vue-awesome/icons/legal'
  import 'vue-awesome/icons/bank'
  import 'vue-awesome/icons/credit-card'
  import 'vue-awesome/icons/ticket'
  import 'vue-awesome/icons/phone'
  import 'vue-awesome/icons/exchange'
  import 'vue-awesome/icons/money'
  import 'vue-awesome/icons/registered'
  import 'vue-awesome/icons/sticky-note'
  import 'vue-awesome/icons/tag'
  import 'vue-awesome/icons/info-circle'
  import {TheMask} from 'vue-the-mask'
  export default {
    name: 'panelModalFranchisesAdd',
    props: ['franchises', 'staff', 'clinics'],
    data () {
      return {
        name: '',
        official_name: '',
        full_official_name: '',
        license: '',
        city: '',
        legal_address: '',
        inn: '',
        kpp: '',
        bank_account: '',
        bank: '',
        correspondent_account: '',
        bik: '',
        ogrn: '',
        okved: '',
        act: '',
        correspondent_address: '',
        number: ''
      }
    },
    methods: {
      addFranchise () {
        this.$emit('addFranchise', {
          name: this.name,
          official_name: this.official_name,
          full_official_name: this.full_official_name,
          license: this.license,
          city: this.city,
          legal_address: this.legal_address,
          inn: this.inn,
          kpp: this.kpp,
          bank_account: this.bank_account,
          bank: this.bank,
          correspondent_account: this.correspondent_account,
          bik: this.bik,
          ogrn: this.ogrn,
          okved: this.okved,
          act: this.act,
          correspondent_address: this.correspondent_address,
          number: this.number
        })
      },
      hideModal () {
        this.$emit('hideModal', {}, true)
      }
    },
    components: {TheMask}
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
        flex-direction: column;
        justify-content: flex-start;
        align-items: center;
        margin: 15px 0 30px 0;
        width: 350px;
    }

</style>
