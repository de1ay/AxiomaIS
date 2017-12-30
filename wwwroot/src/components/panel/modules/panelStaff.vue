<template>
  <div id="panel_media">
    <b>Раздел находится в переработке, релиз 1-2 января.</b>
    <div class="media">
      <vue-good-table
        :columns="media_columns"
        :rows="media"
        :paginate="true"
        :lineNumbers="true"
        :rowsPerPageText="rowsPerPage"
        :nextText="nextText"
        :prevText="prevText"
        :onClick="showMedia"
        :ofText="ofText">
        <div slot="emptystate">Медиа-носители отсутствуют</div>
      </vue-good-table>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'panelStaff',
    props: ['franchises', 'media', 'clients', 'deals'],
    data () {
      return {
        ofText: 'из',
        nextText: 'След.',
        prevText: 'Пред.',
        rowsPerPage: 'Показывать строк:',
        media_columns: [
          {
            label: 'Название',
            field: 'media_name'
          },
          {
            label: 'Адрес',
            field: 'media_address'
          },
          {
            label: 'Тип',
            field: this.parseMediaType
          }
        ]
      }
    },
    methods: {
      parseMediaType (rowObj) {
        switch (rowObj.media_type) {
          case '1': return 'Баннер'
          case '2': return 'Радиостанция'
          default: return 'Ошибка'
        }
      },
      showMedia (rowObj, index) {
        this.$emit('showModal', 'panel_media_view', rowObj)
      }
    }
  }
</script>

<style>

  #panel_media { margin:50px; }

  .media .table-bordered th, .media .table.table-bordered td {
    border-top: none !important;
    border-right: none !important;
    border-left: none !important;
    border-bottom: 0.5px rgba(52, 152, 219, 0.4) solid !important;
  }

  .media .table.condensed thead th, .media .table thead th {
    background-color: #3498db !important;
    color: #fff;
  }

  .media .good-table .table-footer {
    background: transparent;
    border: none;
  }

  .media .good-table .table-footer .pagination-controls {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
  }

  .media .table .sorting-asc, .media .table .sorting-desc {
    color: #fff !important;
    user-select: none;
  }

  .media .table .sorting-asc:after, .media .table .sorting-desc:after {
    border-top-color: #fff !important;
    border-bottom-color: #fff !important;
  }

  .media .good-table .table tbody th.line-numbers {
    background-color: transparent !important;
  }

  .media .good-table .table-footer .pagination-controls .info { font-size: 14px; }

</style>
