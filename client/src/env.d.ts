/// <reference types="vite/client" />

interface ImportMetaEnv {
    readonly VITE_WYSIWYG_REDACTOR_API_KEY?: string;
}

interface ImportMeta {
    readonly env: ImportMetaEnv
}