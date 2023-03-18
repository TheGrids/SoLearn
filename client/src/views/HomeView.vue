<template>
   <div class="container">
        <div class="l-wrp">
            <div class="landing-top__left">
                <div class="name">Повышай уровень знаний</div>
                <div class="second-name">Помогай получать знания другим</div>
                <p class="p">Наш сервис вместил в себе и то и другое</p>

                <Button label="Попробовать" type="button"></Button>
            </div>
            <div class="landing-top__right">
                <img src="https://clubtraining.com.au/wp-content/uploads/2021/09/shutterstock_1214042104-LR-scaled.jpg" alt="Image" class="img">
            </div>
        </div>
    </div>

    <section class="section_2">
        <div class="container">
            <div class="sec-wrp">
                <div class="section_left">
                    <div class="img-wrp">
                        <img src="http://dummyimage.com/640x360" class="img"/>
                    </div>
                </div>
                <div class="section_right">
                    <div class="section-info__name">Какой-то заголовок</div>
                    <div class="section-info__dop">Lorem ipsum dolor, sit amet consectetur adipisicing elit. Corporis quos officia fugiat officiis repellat id, aliquid inventore dolore, iure molestias eveniet doloremque, voluptatum voluptatibus minima distinctio beatae est vitae! Pariatur?</div>
                </div>
            </div>
        </div>
    </section>

    <section class="section_1">
        <div class="container">
            <div class="sec-wrp">
                <div class="section_left">
                    <div class="section-info__name">Какой-то заголовок</div>
                    <div class="section-info__dop">Lorem ipsum dolor, sit amet consectetur adipisicing elit. Corporis quos officia fugiat officiis repellat id, aliquid inventore dolore, iure molestias eveniet doloremque, voluptatum voluptatibus minima distinctio beatae est vitae! Pariatur?</div>
                </div>
                <div class="section_right">
                    <div class="img-wrp">
                        <img src="http://dummyimage.com/640x360" class="img"/>
                    </div>
                </div>
            </div>
        </div>
    </section>

    <section class="section_2">
        <div class="result">
            <div class="result-main">Присоединяйся</div>
            <div class="p">Начинай повышать уровень знаний</div>
            <Button label="Попробовать" type="button"></Button>
        </div>
    </section>

    <footer>
        <div class="container">
            <div class="footer">
                <div class="footer-logo">
                    <span style="font-weight: bold;color: #383B41;">So</span><span style="font-weight: bold;color: #697187">Learn</span>
                </div>
                <div class="footer-info">
                    <div class="footer-info__item">
                        <div class="item-name">Сообщество</div>
                        <div class="item-list">github</div>
                        <div class="item-list">vk</div>
                        <div class="item-list">github</div>
                    </div>

                    <div class="footer-info__item">
                        <div class="item-name">Сервис</div>
                        <div class="item-list">О нас</div>
                        <div class="item-list">Новости</div>
                        <div class="item-list">Вакансии</div>
                    </div>
                </div>
            </div>
        </div>
    </footer>

</template>
<script>
export default {
  name: 'app'
}
</script>

<style lang="scss" scoped>
.footer {
    padding: 50px 20px;
    width: 100%;
    color: #383B41;
    display: flex;
    justify-content: space-between;

    .footer-logo {
        font-size: 32px;
    }

    .footer-info {
        display: flex;

        .footer-info__item {
            margin-left: 200px;

            .item-name {
                font-size: 24px;
                font-weight: 700;
            }
            .item-list {
                font-size: 20px;
            }
        }
    }
}

.section_1 {
    padding: 20px 0;
    color: #383B41;

    .sec-wrp {
        display: flex;
        align-items: center;
        justify-content: space-between;
    
        .img-wrp {
            min-width: 360px;

            .img {
                width: 100%;
                height: 100%;
                border-radius: 25px;
            }
        }
        .section_left {
            width: 600px;
        }
    }
}
.section_2 {
    background: #EAEAEA;
    padding: 20px 0;
    color: #383B41;

    .sec-wrp {
        display: flex;
        align-items: center;
    
        .img-wrp {
            min-width: 360px;

            .img {
                width: 100%;
                height: 100%;
                border-radius: 25px;
            }
        }
        .section_right {
            width: 600px;
        }
    }
}
.container {
    max-width: 1440px;
    margin: 0 auto;
    padding: 0 20px;
}
.l-wrp {
    display: flex;
    justify-content: space-between;
    align-items: center;
}
.landing-top__left {
    color: #383B41;
    min-width: 300px;

    .name {
        font-size: 38px;
        font-weight: 700;
    }
    .second-name {
        font-size: 38px;
        color: #697187
    }
    .p {
        font-size: 20px;
        margin-bottom: 20px;
    }
}
.landing-top__right {
    max-width: 700px;

    .img {
        width: 100%;
        clip-path: polygon(8% 0, 100% 0%, 100% 100%, 0 100%)
    }
}


.section_left {
    padding-right: 20px;
}
.section-info__name {
    font-size: 38px;
    font-weight: 700;
}
.section-info__dop {
    font-size: 20px;
}
.result {
    text-align: center;
    padding: 50px 0;
}
.result-main {
    font-size: 38px;
    font-weight: 700;
}
.p {
    font-size: 24px;
    margin-bottom: 20px;
}
@media screen  and (max-width: 768px) {  
    .l-wrp {
        display: block;
        text-align: center;
    }
    .landing-top__right {
        margin-top: 20px;
        max-width: 100%;
        .img {
            clip-path: none!important;
        }
    }

    .section_1 {
        .sec-wrp {
            display: block;

            .section_right {
                margin-top: 20px;
                padding: 0;
            }
            .section_left {
                margin: 0;
                width: 100%;
            }
        }
    }

    .section_2 {
        .sec-wrp {
            display: block;

            .section_left {
                padding: 0;
            }
            .section_right {
                margin: 0;
                width: 100%;
            }
        }
    }

    .footer-info__item {
        margin-left: 20px!important;
    }
}
</style>