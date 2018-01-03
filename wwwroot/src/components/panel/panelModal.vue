<template>
  <div id="panel_modal" @click="hideModal">
    <div class="modal">
      <button class="panel_modal-button-hide" @click="hideModal({}, true)">Закрыть</button>
      <component
        :is="modal_name"
        :franchises.sync="franchises"
        :clinics.sync="clinics"
        :staff.sync="staff"
        :additional_data="additional_data"
        @hideModal="hideModal"
        @addDeal="addDeal"
        @editDeal="editDeal"
        @deleteDeal="deleteDeal"
        @addClinic="addClinic"
        @editClinic="editClinic"
        @deleteClinic="deleteClinic"
        @addFranchise="addFranchise"
        @editFranchise="editFranchise"
        @deleteFranchise="deleteFranchise">
      </component>
    </div>
  </div>
</template>

<script>
  import axios from 'axios'
  import panelModalStaffView from '@/components/panel/modals/panelModalStaffView'
  import panelModalStaffAdd from '@/components/panel/modals/panelModalStaffAdd'
  import panelModalClinicsView from '@/components/panel/modals/panelModalClinicsView'
  import panelModalClinicsAdd from '@/components/panel/modals/panelModalClinicsAdd'
  import panelModalFranchisesView from '@/components/panel/modals/panelModalFranchisesView'
  import panelModalFranchisesAdd from '@/components/panel/modals/panelModalFranchisesAdd'
  export default {
    name: 'panelModal',
    props: ['franchises', 'modal_name', 'modal_active', 'staff', 'clinics', 'additional_data', 'token'],
    components: {
      'panel_staff_view': panelModalStaffView,
      'panel_staff_add': panelModalStaffAdd,
      'panel_clinics_view': panelModalClinicsView,
      'panel_clinics_add': panelModalClinicsAdd,
      'panel_franchises_add': panelModalFranchisesAdd,
      'panel_franchises_view': panelModalFranchisesView
    },
    methods: {
      showModal (modalName) {
        this.$emit('showModal', modalName)
      },
      hideModal (e, flag = false) {
        if (flag || e.target.id === 'panel_modal') {
          this.$emit('update:modal_active', false)
        }
      },
      addDeal (deal) {
        this.$snotify.async(
          'Запрос выполняется',
          'Подождите...',
          () => new Promise((resolve, reject) => {
            axios.post('https://beta.project.nullteam.info/api/deals/', {
              deal_client: deal.deal_client,
              deal_brand: deal.deal_brand,
              deal_media: deal.deal_media,
              deal_period: deal.deal_period,
              deal_time: deal.deal_time,
              deal_status: deal.deal_status,
              deal_sum: deal.deal_sum,
              deal_paid: deal.deal_paid,
              deal_type: deal.deal_type
            }).then(resp => {
              deal.deal_id = resp.data.deal_id
              this.deals.push(deal)
              resolve({
                title: 'Успешно',
                body: 'Сделка добавлена',
                config: {
                  closeOnClick: true,
                  timeout: 2000,
                  showProgressBar: true,
                  pauseOnHover: true
                }
              })
            }).catch(resp => {
              /*eslint-disable */
              reject({
                title: 'Ошибка!',
                body: 'Сделка не добавлена',
                config: {
                  closeOnClick: true,
                  timeout: 2000,
                  showProgressBar: true,
                  pauseOnHover: true
                }
              })
              /*eslint-enable */
            })
          }
        ))
        this.$emit('update:deals', this.deals)
        this.$emit('update:modal_active', false)
      },
      editDeal (newDeal, originalDeal) {
        this.$snotify.async(
          'Запрос выполняется',
          'Подождите...',
          () => new Promise((resolve, reject) => {
            axios.put('https://beta.project.nullteam.info/api/deals/' + newDeal.deal_id, {
              deal_client: newDeal.deal_client,
              deal_brand: newDeal.deal_brand,
              deal_media: newDeal.deal_media,
              deal_period: newDeal.deal_period,
              deal_time: newDeal.deal_time,
              deal_type: newDeal.deal_type,
              deal_status: newDeal.deal_status,
              deal_sum: newDeal.deal_sum,
              deal_paid: newDeal.deal_paid
            }).then(resp => {
              originalDeal = Object.assign(originalDeal, newDeal)
              resolve({
                title: 'Успешно',
                body: 'Сделка изменена',
                config: {
                  closeOnClick: true,
                  timeout: 2000,
                  showProgressBar: true,
                  pauseOnHover: true
                }
              })
            }).catch(resp => {
              /*eslint-disable */
              reject({
                title: 'Ошибка!',
                body: 'Сделка не изменена',
                config: {
                  closeOnClick: true,
                  timeout: 2000,
                  showProgressBar: true,
                  pauseOnHover: true
                }
              })
              /*eslint-enable */
            })
          }
        ))
        this.$emit('update:modal_active', false)
      },
      deleteDeal (deal) {
        this.$snotify.async(
          'Запрос выполняется',
          'Подождите...',
          () => new Promise((resolve, reject) => {
            axios.delete('https://beta.project.nullteam.info/api/deals/' + deal.deal_id).then(resp => {
              let relatedBillings = this.franchises.filter(billing => {
                return billing.billing_deal_info.deal_id === deal.deal_id
              })
              relatedBillings.forEach(billing => {
                this.franchises.splice(this.franchises.indexOf(billing))
              })
              this.deals.splice(this.deals.indexOf(deal))
              this.$emit('update:deals', this.deals)
              this.$emit('update:franchises', this.franchises)
              resolve({
                title: 'Успешно',
                body: 'Сделка удалена',
                config: {
                  closeOnClick: true,
                  timeout: 2000,
                  showProgressBar: true,
                  pauseOnHover: true
                }
              })
            }).catch(resp => {
              /*eslint-disable */
              reject({
                title: 'Ошибка!',
                body: 'Сделка не удалена',
                config: {
                  closeOnClick: true,
                  timeout: 2000,
                  showProgressBar: true,
                  pauseOnHover: true
                }
              })
              /*eslint-enable */
            })
          }
        ))
        this.hideModal({}, true)
      },
      addClinic (clinic) {
        this.$snotify.async(
          'Запрос выполняется',
          'Подождите...',
          () => new Promise((resolve, reject) => {
            axios.post('https://api.axiomais.ru/clinics/', {
              token: this.token,
              name: clinic.name,
              address: clinic.address,
              is_active: clinic.is_active,
              city: clinic.city,
              franchise_id: clinic.franchise_id
            }).then(resp => {
              if (resp.data.code !== 200) {
                /*eslint-disable */
                reject({
                  title: 'Ошибка!',
                  body: resp.data.message,
                  config: {
                    closeOnClick: true,
                    timeout: 2000,
                    showProgressBar: true,
                    pauseOnHover: true
                  }
                })
                /*eslint-enable */
              } else {
                clinic.id = resp.data.message.id
                this.clinics.push(clinic)
                this.$emit('update:clinics', this.clinics)
                resolve({
                  title: 'Успешно',
                  body: 'Клиника добавлена',
                  config: {
                    closeOnClick: true,
                    timeout: 2000,
                    showProgressBar: true,
                    pauseOnHover: true
                  }
                })
              }
            }).catch(resp => {
              /*eslint-disable */
              reject({
                title: 'Ошибка!',
                body: resp.data.message,
                config: {
                  closeOnClick: true,
                  timeout: 2000,
                  showProgressBar: true,
                  pauseOnHover: true
                }
              })
              /*eslint-enable */
            })
          }
        ))
        this.$emit('update:modal_active', false)
      },
      editClinic (newClinic, originalClinic) {
        this.$snotify.async(
          'Запрос выполняется',
          'Подождите...',
          () => new Promise((resolve, reject) => {
            axios.put('https://api.axiomais.ru/clinics/' + newClinic.id, {
              token: this.token,
              id: newClinic.id,
              name: newClinic.name,
              address: newClinic.address,
              is_active: newClinic.is_active,
              city: newClinic.city,
              franchise_id: newClinic.franchise_id
            }).then(resp => {
              if (resp.data.code !== 200) {
                /*eslint-disable */
                reject({
                  title: 'Ошибка!',
                  body: resp.data.message,
                  config: {
                    closeOnClick: true,
                    timeout: 2000,
                    showProgressBar: true,
                    pauseOnHover: true
                  }
                })
                /*eslint-enable */
              } else {
                originalClinic = Object.assign(originalClinic, newClinic)
                this.$emit('update:clinics', this.clinics)
                resolve({
                  title: 'Успешно',
                  body: 'Клиника изменена',
                  config: {
                    closeOnClick: true,
                    timeout: 2000,
                    showProgressBar: true,
                    pauseOnHover: true
                  }
                })
              }
            }).catch(resp => {
              /*eslint-disable */
              reject({
                title: 'Ошибка!',
                body: resp.data.message,
                config: {
                  closeOnClick: true,
                  timeout: 2000,
                  showProgressBar: true,
                  pauseOnHover: true
                }
              })
              /*eslint-enable */
            })
          }
        ))
        this.$emit('update:modal_active', false)
      },
      deleteClinic (clinic) {
        this.$snotify.async(
          'Запрос выполняется',
          'Подождите...',
          () => new Promise((resolve, reject) => {
            axios.delete('https://api.axiomais.ru/clinics/' + clinic.id, {
              token: this.token,
              id: clinic.id,
              name: clinic.name,
              address: clinic.address,
              is_active: clinic.is_active,
              city: clinic.city,
              franchise_id: clinic.franchise_id
            }).then(resp => {
              if (resp.data.code !== 200) {
                /*eslint-disable */
                reject({
                  title: 'Ошибка!',
                  body: resp.data.message,
                  config: {
                    closeOnClick: true,
                    timeout: 2000,
                    showProgressBar: true,
                    pauseOnHover: true
                  }
                })
                /*eslint-enable */
              } else {
                this.clinics.splice(this.clinics.indexOf(clinic))
                this.$emit('update:clinics', this.clinics)
                resolve({
                  title: 'Успешно',
                  body: 'Клиника удалена',
                  config: {
                    closeOnClick: true,
                    timeout: 2000,
                    showProgressBar: true,
                    pauseOnHover: true
                  }
                })
              }
            }).catch(resp => {
              /*eslint-disable */
              reject({
                title: 'Ошибка!',
                body: 'Клиника не удалена',
                config: {
                  closeOnClick: true,
                  timeout: 2000,
                  showProgressBar: true,
                  pauseOnHover: true
                }
              })
              /*eslint-enable */
            })
          }
        ))
        this.hideModal({}, true)
      },
      addBilling (billing) {
        this.$snotify.async(
          'Запрос выполняется',
          'Подождите...',
          () => new Promise((resolve, reject) => {
            axios.post('https://beta.project.nullteam.info/api/franchises/', {
              billing_deal: billing.billing_deal,
              billing_sum: billing.billing_sum,
              billing_date: billing.billing_date,
              billing_transfer_date: billing.billing_transfer_date
            }).then(resp => {
              this.franchises.push({
                billing_id: resp.data.billing_id,
                billing_sum: resp.data.billing_sum,
                billing_date: resp.data.billing_date,
                billing_transfer_date: billing.billing_transfer_date,
                billing_deal: resp.data.billing_deal,
                billing_deal_info: billing.billing_deal_info
              })
              this.$emit('update:franchises', this.franchises)
              let relatedDeal = this.deals.filter(deal => {
                return deal.deal_id === billing.billing_deal
              })[0]
              relatedDeal.deal_paid += billing.billing_sum
              if (relatedDeal.deal_paid >= relatedDeal.deal_sum) {
                relatedDeal.deal_status = '2'
              } else {
                relatedDeal.deal_status = '1'
              }
              this.$emit('update:deals', this.deals)
              resolve({
                title: 'Успешно',
                body: 'Оплата добавлена',
                config: {
                  closeOnClick: true,
                  timeout: 2000,
                  showProgressBar: true,
                  pauseOnHover: true
                }
              })
            }).catch(resp => {
              /*eslint-disable */
              reject({
                title: 'Ошибка!',
                body: 'Оплата не добавлена',
                config: {
                  closeOnClick: true,
                  timeout: 2000,
                  showProgressBar: true,
                  pauseOnHover: true
                }
              })
              /*eslint-enable */
            })
          }
        ))
        this.$emit('update:modal_active', false)
      },
      editBilling (newBilling, originalBilling) {
        this.$snotify.async(
          'Запрос выполняется',
          'Подождите...',
          () => new Promise((resolve, reject) => {
            axios.put('https://beta.project.nullteam.info/api/franchises/' + newBilling.billing_id, {
              billing_deal: newBilling.billing_deal,
              billing_date: newBilling.billing_date,
              billing_sum: newBilling.billing_sum,
              billing_transfer_date: newBilling.billing_transfer_date
            }).then(resp => {
              if (newBilling.billing_deal === originalBilling.billing_deal) {
                let relatedDeal = this.deals.filter(deal => {
                  return deal.deal_id === originalBilling.billing_deal
                })[0]
                relatedDeal.deal_paid = relatedDeal.deal_paid - originalBilling.billing_sum + newBilling.billing_sum
                if (relatedDeal.deal_paid >= relatedDeal.deal_sum) {
                  relatedDeal.deal_status = '2'
                } else {
                  relatedDeal.deal_status = '1'
                }
              } else {
                let originalRelatedDeal = this.deals.filter(deal => {
                  return deal.deal_id === originalBilling.billing_deal
                })[0]
                let newRelatedDeal = this.deals.filter(deal => {
                  return deal.deal_id === newBilling.billing_deal
                })[0]
                originalRelatedDeal.deal_paid -= originalBilling.billing_sum
                if (originalRelatedDeal.deal_paid >= originalRelatedDeal.deal_sum) {
                  originalRelatedDeal.deal_status = '2'
                } else {
                  originalRelatedDeal.deal_status = '1'
                }
                newRelatedDeal.deal_paid += newBilling.billing_sum
                if (newRelatedDeal.deal_paid >= newRelatedDeal.deal_sum) {
                  newRelatedDeal.deal_status = '2'
                } else {
                  newRelatedDeal.deal_status = '1'
                }
              }
              this.franchises.splice(this.franchises.indexOf(originalBilling))
              this.franchises.push(newBilling)
              this.$emit('update:deals', this.deals)
              this.$emit('update:franchises', this.franchises)
              resolve({
                title: 'Успешно',
                body: 'Оплата изменена',
                config: {
                  closeOnClick: true,
                  timeout: 2000,
                  showProgressBar: true,
                  pauseOnHover: true
                }
              })
            }).catch(resp => {
              /*eslint-disable */
              reject({
                title: 'Ошибка!',
                body: 'Оплата не изменена',
                config: {
                  closeOnClick: true,
                  timeout: 2000,
                  showProgressBar: true,
                  pauseOnHover: true
                }
              })
              /*eslint-enable */
            })
          }
        ))
        this.$emit('update:modal_active', false)
      },
      deleteBilling (billing) {
        this.$snotify.async(
          'Запрос выполняется',
          'Подождите...',
          () => new Promise((resolve, reject) => {
            axios.delete('https://beta.project.nullteam.info/api/franchises/' + billing.billing_id).then(resp => {
              let relatedDeal = this.deals.filter(deal => {
                return deal.deal_id === billing.billing_deal_info.deal_id
              })[0]
              relatedDeal.deal_paid -= billing.billing_sum
              if (relatedDeal.deal_paid < relatedDeal.deal_sum) {
                relatedDeal.deal_status = '1'
              }
              this.franchises.splice(this.franchises.indexOf(billing))
              this.$emit('update:franchises', this.franchises)
              resolve({
                title: 'Успешно',
                body: 'Оплата удалена',
                config: {
                  closeOnClick: true,
                  timeout: 2000,
                  showProgressBar: true,
                  pauseOnHover: true
                }
              })
            }).catch(resp => {
              /*eslint-disable */
              reject({
                title: 'Ошибка!',
                body: 'Оплата не удалена',
                config: {
                  closeOnClick: true,
                  timeout: 2000,
                  showProgressBar: true,
                  pauseOnHover: true
                }
              })
              /*eslint-enable */
            })
          }
        ))
        this.hideModal({}, true)
      },
      addFranchise (franchise) {
        this.$snotify.async(
          'Запрос выполняется',
          'Подождите...',
          () => new Promise((resolve, reject) => {
            axios.post('https://api.axiomais.ru/franchises', {
              token: this.token,
              name: franchise.name,
              official_name: franchise.official_name,
              full_official_name: franchise.full_official_name,
              license: franchise.license,
              city: franchise.city,
              legal_address: franchise.legal_address,
              inn: franchise.inn,
              kpp: franchise.kpp,
              bank_account: franchise.bank_account,
              bank: franchise.bank,
              correspondent_account: franchise.correspondent_account,
              bik: franchise.bik,
              ogrn: franchise.ogrn,
              okved: franchise.okved,
              act: franchise.act,
              correspondent_address: franchise.correspondent_address,
              number: franchise.number
            }).then(resp => {
              if (resp.data.code !== 200) {
                /*eslint-disable */
                reject({
                  title: 'Ошибка!',
                  body: resp.data.message,
                  config: {
                    closeOnClick: true,
                    timeout: 2000,
                    showProgressBar: true,
                    pauseOnHover: true
                  }
                })
                /*eslint-enable */
              } else {
                this.franchises.push(resp.data.message)
                this.$emit('update:franchises', this.franchises)
                resolve({
                  title: 'Успешно',
                  body: 'Франчайзи добавлен',
                  config: {
                    closeOnClick: true,
                    timeout: 2000,
                    showProgressBar: true,
                    pauseOnHover: true
                  }
                })
              }
            }).catch(resp => {
              /*eslint-disable */
              reject({
                title: 'Ошибка!',
                body: resp.data.message,
                config: {
                  closeOnClick: true,
                  timeout: 2000,
                  showProgressBar: true,
                  pauseOnHover: true
                }
              })
              /*eslint-enable */
            })
          }
        ))
        this.$emit('update:modal_active', false)
      },
      editFranchise (newFranchise, originalFranchise) {
        this.$snotify.async(
          'Запрос выполняется',
          'Подождите...',
          () => new Promise((resolve, reject) => {
            axios.put('https://api.axiomais.ru/franchises/' + newFranchise.id, {
              token: this.token,
              id: newFranchise.id,
              name: newFranchise.name,
              official_name: newFranchise.official_name,
              full_official_name: newFranchise.full_official_name,
              license: newFranchise.license,
              city: newFranchise.city,
              legal_address: newFranchise.legal_address,
              inn: newFranchise.inn,
              kpp: newFranchise.kpp,
              bank_account: newFranchise.bank_account,
              bank: newFranchise.bank,
              correspondent_account: newFranchise.correspondent_account,
              bik: newFranchise.bik,
              ogrn: newFranchise.ogrn,
              okved: newFranchise.okved,
              act: newFranchise.act,
              correspondent_address: newFranchise.correspondent_address,
              number: newFranchise.number
            }).then(resp => {
              if (resp.data.code !== 200) {
                /*eslint-disable */
                reject({
                  title: 'Ошибка!',
                  body: resp.data.message,
                  config: {
                    closeOnClick: true,
                    timeout: 2000,
                    showProgressBar: true,
                    pauseOnHover: true
                  }
                })
                /*eslint-enable */
              } else {
                originalFranchise = Object.assign(originalFranchise, newFranchise)
                resolve({
                  title: 'Успешно',
                  body: 'Франчайзи изменёна',
                  config: {
                    closeOnClick: true,
                    timeout: 2000,
                    showProgressBar: true,
                    pauseOnHover: true
                  }
                })
              }
            }).catch(resp => {
              /*eslint-disable */
              reject({
                title: 'Ошибка!',
                body: resp.data.message,
                config: {
                  closeOnClick: true,
                  timeout: 2000,
                  showProgressBar: true,
                  pauseOnHover: true
                }
              })
              /*eslint-enable */
            })
          }
        ))
        this.$emit('update:modal_active', false)
      },
      deleteFranchise (franchise) {
        this.$snotify.async(
          'Запрос выполняется',
          'Подождите...',
          () => new Promise((resolve, reject) => {
            axios.delete('https://api.axiomais.ru/franchises/' + franchise.id, {
              data: {
                token: this.token,
                id: franchise.id,
                name: franchise.name,
                official_name: franchise.official_name,
                full_official_name: franchise.full_official_name,
                license: franchise.license,
                city: franchise.city,
                legal_address: franchise.legal_address,
                inn: franchise.inn,
                kpp: franchise.kpp,
                bank_account: franchise.bank_account,
                bank: franchise.bank,
                correspondent_account: franchise.correspondent_account,
                bik: franchise.bik,
                ogrn: franchise.ogrn,
                okved: franchise.okved,
                act: franchise.act,
                correspondent_address: franchise.correspondent_address,
                number: franchise.number
              }
            }).then(resp => {
              if (resp.data.code !== 200) {
                /*eslint-disable */
                reject({
                  title: 'Ошибка!',
                  body: resp.data.message,
                  config: {
                    closeOnClick: true,
                    timeout: 2000,
                    showProgressBar: true,
                    pauseOnHover: true
                  }
                })
                /*eslint-enable */
              } else {
                this.franchises.splice(this.franchises.indexOf(franchise))
                this.$emit('update:franchises', this.franchises)
                resolve({
                  title: 'Успешно',
                  body: 'Франчайзи удалена',
                  config: {
                    closeOnClick: true,
                    timeout: 2000,
                    showProgressBar: true,
                    pauseOnHover: true
                  }
                })
              }
            }).catch(resp => {
              /*eslint-disable */
              reject({
                title: 'Ошибка!',
                body: resp.data.message,
                config: {
                  closeOnClick: true,
                  timeout: 2000,
                  showProgressBar: true,
                  pauseOnHover: true
                }
              })
              /*eslint-enable */
            })
          }
        ))
        this.hideModal({}, true)
      }
    }
  }
</script>

<style>

    #panel_modal {
      z-index: 1000;
      position: fixed;
      display: flex;
      flex-direction: row;
      justify-content: center;
      align-items: flex-start;
      height: 100%;
      width: 100%;
      background: rgba(0, 0, 0, 0.3);
      overflow-y: auto;
    }

    .modal {
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      margin: 100px 0;
      min-width: 400px;
      min-height: 200px;
      border-radius: 5px;
      background: #ecf0f1;
    }

    #panel_modal .modal .panel_modal-button-hide {
      margin-top: 30px;
      width: 350px;
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

    #panel_modal .modal .panel_modal-button-hide:hover { background: #c0392b; }

    #panel_modal .form-field {
      display: flex;
      flex-direction: row;
      justify-content: center;
      align-items: center;
      margin: 15px 0;
      width: 350px;
      height: 38px;
      border-radius: 5px;
      background-color: #ffffff;
    }

    #panel_modal .form-field--border {
      padding: 0 25px;
      width: 100%;
      border-bottom: 1px solid rgba(0,0,0,.1)
    }

    #panel_modal .form-field:first-child {
        margin: 30px 0 15px 0;
    }

    #panel_modal .form-field .fa-icon { color: #95a5a6; }

    #panel_modal .form-field .form-field__icon--lock {
        margin-top: 4px;
    }

    #panel_modal .form-field .form-field__input {
      margin-right: 10px;
        width: 290px;
        height: 35px;
        font-size: 18px;
        color: #95a5a6;
        border: none;
        outline: none;
    }

    #panel_modal .form-field .form-field__input::-webkit-input-placeholder { color: #95a5a6; }
    #panel_modal .form-field .form-field__input:-moz-placeholder { color: #95a5a6; }
    #panel_modal .form-field .form-field__input::-moz-placeholder { color: #95a5a6; }
    #panel_modal .form-field .form-field__input:-ms-input-placeholder { color: #95a5a6; }

    #panel_modal .form-field__input .multiselect__tags {
        display: flex;
        flex-direction: row;
        justify-content: flex-start;
        align-items: center;
        padding: 0;
        margin: 2px 0;
        min-height: 35px;
        max-height: 35px;
        border: none;
    }

    #panel_modal .form-field__input .multiselect__tags input {
        margin: 0;
        padding: 0;
        height: 35px;
        font-size: 18px;
        color: #95a5a6;
    }

    #panel_modal .form-field__input .multiselect__tags input::-webkit-input-placeholder { color: #95a5a6; }
    #panel_modal .form-field__input .multiselect__tags input:-moz-placeholder { color: #95a5a6; }
    #panel_modal .form-field__input .multiselect__tags input::-moz-placeholder { color: #95a5a6; }
    #panel_modal .form-field__input .multiselect__tags input:-ms-input-placeholder { color: #95a5a6; }

</style>
