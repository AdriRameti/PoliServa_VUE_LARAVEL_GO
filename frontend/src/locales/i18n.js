import { createI18n } from "vue-i18n";
import { useStore } from 'vuex';
import en from './en.json';
import es from './es.json';
import val from './val.json';

const i18n = createI18n({
    locale: 'en',
    messages: {
        en: en,
        es: es,
        val: val
    }
})

export default i18n;