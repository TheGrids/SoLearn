<template>
    <div class="container">
        <div class="lesson-nav__header-m">
            <div class="name">Название курса</div>
            <div class="progress">Прогресс: 1/100</div>
        </div>
    </div>
    <div class="container">
        <div class="lesson-nav">
            <div class="lesson-nav__header">
                <div class="name">Название курса</div>
                <div class="progress">Прогресс: 1/100</div>
            </div>
            <div class="lesson-nav__lessons">
                <div class="item">
                    <div class="item-name">1. <span style="font-weight: bold">Название темы</span></div>
                    <div class="item-subname active">1.1 <span>Название подтемы</span></div>
                    <div class="item-subname">1.1 <span>Название подтемы</span></div>
                    <div class="item-subname">1.1 <span>Название подтемы</span></div>
                </div>

                <div class="item">
                    <div class="item-name">1. <span style="font-weight: bold">Название темы</span></div>
                    <div class="item-subname">1.1 <span>Название подтемы</span></div>
                    <div class="item-subname">1.1 <span>Название подтемы</span></div>
                    <div class="item-subname">1.1 <span>Название подтемы</span></div>
                </div>

            </div>
        </div>
        <div class="lesson-main">
            <div class="name-lesson">1. Название темы <span style="color: #D9D9D9">/</span> <span>1.1 Название подтемы</span></div>

            <img src="http://dummyimage.com/1200x500" class="ia"/>
            <Button style="margin: 10px 10px 0 0" icon="pi pi-angle-left" label="Предыдущий шаг" severity="secondary" rounded />
            <Button style="margin: 10px 10px 0 0" icon="pi pi-angle-right" iconPos="right" label="Следующий шаг" severity="secondary" rounded />
        </div>
    </div>
</template>

<script>
export default {
    name: 'lesson'
}
</script>

<style lang="scss" scoped>
.container {
    display: flex;
    max-width: 1440px;
    margin: 0 auto;
    padding: 0 20px;
    color: #383B41;
}

.lesson-nav__header-m {
    display: none;
    margin-top: 20px;
    margin-bottom: 20px;
    background: #e3e3e3;
    border-radius: 25px;
    text-align: center;
    padding: 10px 0;
    width: 100%;

    .name {
        font-weight: 700;
        font-size: 24px;
    }
    .progress {
        font-weight: 400;
        font-size: 24px;
    }
}

.lesson-nav {
    min-width: 300px;
    margin-right: 20px;


    .lesson-nav__header {
        background: #e3e3e3;
        border-radius: 25px;
        text-align: center;
        padding: 10px 0;
        width: 100%;

        .name {
            font-weight: 700;
            font-size: 24px;
        }
        .progress {
            font-weight: 400;
            font-size: 24px;
        }
    }
    .lesson-nav__lessons {
        padding: 25px 0;

        .item {


            .item-name {
                font-weight: 700;
                font-size: 24px;
            }
            .item-subname {
                font-weight: 400;
                font-size: 24px;
                padding: 5px 20px 5px 20px;
                border-radius: 25px;
            }
            .active {
                background:#e3e3e3;
            }
        }
    }
}
.lesson-main {
    .ia {
        width: 100%;
        
        .next-btn {
            float: right;
        }
    }
    .name-lesson {
        font-weight: 400;
        font-size: 24px;
        font-weight: bold;
        margin: 20px 0;

        span {

        }
    }
}
@media screen  and (max-width: 1024px) {  
    .item-name span {
        display: none;
    }
    .item-subname span {
        display: none;
    }
    .lesson-nav {
        min-width: auto;
    }
    .lesson-nav__header {
        display: none;
    }
    .lesson-nav__header-m {
        display: block;
    }
}
</style>