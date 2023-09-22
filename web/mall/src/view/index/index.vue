<template>
  <div>
    <!--轮播图-->
    <div style="width: 1220px;margin: 0 auto" class="mall_carousel_category" @mouseleave="is_hover_category = false">
      <el-carousel :interval="2000" arrow="never" indicator-position="none" style="height: 100%;width: 100%">
        <el-carousel-item v-for="item in carouselItemList" :key="item">
          <el-image style="width: 100%;height: 100%;" :src="item.url" />
        </el-carousel-item>
      </el-carousel>
      <!--侧面分类-->
      <div class="mall_category">
        <div @mouseover="is_hover_category = true">
          <p>手机</p>
          <svg
            t="1694692637616"
            class="icon"
            viewBox="0 0 1024 1024"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            p-id="3973"
            width="200"
            height="200"
          >
            <path
              d="M761.055557 532.128047c0.512619-0.992555 1.343475-1.823411 1.792447-2.848649 8.800538-18.304636 5.919204-40.703346-9.664077-55.424808L399.935923 139.743798c-19.264507-18.208305-49.631179-17.344765-67.872168 1.888778-18.208305 19.264507-17.375729 49.631179 1.888778 67.872168l316.960409 299.839269L335.199677 813.631716c-19.071845 18.399247-19.648112 48.767639-1.247144 67.872168 9.407768 9.791372 21.984142 14.688778 34.560516 14.688778 12.000108 0 24.000215-4.479398 33.311652-13.439914l350.048434-337.375729c0.672598-0.672598 0.927187-1.599785 1.599785-2.303346 0.512619-0.479935 1.056202-0.832576 1.567101-1.343475C757.759656 538.879828 759.199462 535.391265 761.055557 532.128047z"
              fill="#ffffff"
              p-id="3974"
            />
          </svg>
        </div>
        <div @mouseover="is_hover_category = true">
          <p>洗衣机</p>
          <svg
            t="1694692637616"
            class="icon"
            viewBox="0 0 1024 1024"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            p-id="3973"
            width="200"
            height="200"
          >
            <path
              d="M761.055557 532.128047c0.512619-0.992555 1.343475-1.823411 1.792447-2.848649 8.800538-18.304636 5.919204-40.703346-9.664077-55.424808L399.935923 139.743798c-19.264507-18.208305-49.631179-17.344765-67.872168 1.888778-18.208305 19.264507-17.375729 49.631179 1.888778 67.872168l316.960409 299.839269L335.199677 813.631716c-19.071845 18.399247-19.648112 48.767639-1.247144 67.872168 9.407768 9.791372 21.984142 14.688778 34.560516 14.688778 12.000108 0 24.000215-4.479398 33.311652-13.439914l350.048434-337.375729c0.672598-0.672598 0.927187-1.599785 1.599785-2.303346 0.512619-0.479935 1.056202-0.832576 1.567101-1.343475C757.759656 538.879828 759.199462 535.391265 761.055557 532.128047z"
              fill="#ffffff"
              p-id="3974"
            />
          </svg>
        </div>
        <div @mouseover="is_hover_category = true">
          <p>家电</p>
          <svg
            t="1694692637616"
            class="icon"
            viewBox="0 0 1024 1024"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            p-id="3973"
            width="200"
            height="200"
          >
            <path
              d="M761.055557 532.128047c0.512619-0.992555 1.343475-1.823411 1.792447-2.848649 8.800538-18.304636 5.919204-40.703346-9.664077-55.424808L399.935923 139.743798c-19.264507-18.208305-49.631179-17.344765-67.872168 1.888778-18.208305 19.264507-17.375729 49.631179 1.888778 67.872168l316.960409 299.839269L335.199677 813.631716c-19.071845 18.399247-19.648112 48.767639-1.247144 67.872168 9.407768 9.791372 21.984142 14.688778 34.560516 14.688778 12.000108 0 24.000215-4.479398 33.311652-13.439914l350.048434-337.375729c0.672598-0.672598 0.927187-1.599785 1.599785-2.303346 0.512619-0.479935 1.056202-0.832576 1.567101-1.343475C757.759656 538.879828 759.199462 535.391265 761.055557 532.128047z"
              fill="#ffffff"
              p-id="3974"
            />
          </svg>
        </div>
        <div @mouseover="is_hover_category = true">
          <p>笔记本</p>
          <svg
            t="1694692637616"
            class="icon"
            viewBox="0 0 1024 1024"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            p-id="3973"
            width="200"
            height="200"
          >
            <path
              d="M761.055557 532.128047c0.512619-0.992555 1.343475-1.823411 1.792447-2.848649 8.800538-18.304636 5.919204-40.703346-9.664077-55.424808L399.935923 139.743798c-19.264507-18.208305-49.631179-17.344765-67.872168 1.888778-18.208305 19.264507-17.375729 49.631179 1.888778 67.872168l316.960409 299.839269L335.199677 813.631716c-19.071845 18.399247-19.648112 48.767639-1.247144 67.872168 9.407768 9.791372 21.984142 14.688778 34.560516 14.688778 12.000108 0 24.000215-4.479398 33.311652-13.439914l350.048434-337.375729c0.672598-0.672598 0.927187-1.599785 1.599785-2.303346 0.512619-0.479935 1.056202-0.832576 1.567101-1.343475C757.759656 538.879828 759.199462 535.391265 761.055557 532.128047z"
              fill="#ffffff"
              p-id="3974"
            />
          </svg>
        </div>
        <div @mouseover="is_hover_category = true">
          <p>出行 穿戴</p>
          <svg
            t="1694692637616"
            class="icon"
            viewBox="0 0 1024 1024"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            p-id="3973"
            width="200"
            height="200"
          >
            <path
              d="M761.055557 532.128047c0.512619-0.992555 1.343475-1.823411 1.792447-2.848649 8.800538-18.304636 5.919204-40.703346-9.664077-55.424808L399.935923 139.743798c-19.264507-18.208305-49.631179-17.344765-67.872168 1.888778-18.208305 19.264507-17.375729 49.631179 1.888778 67.872168l316.960409 299.839269L335.199677 813.631716c-19.071845 18.399247-19.648112 48.767639-1.247144 67.872168 9.407768 9.791372 21.984142 14.688778 34.560516 14.688778 12.000108 0 24.000215-4.479398 33.311652-13.439914l350.048434-337.375729c0.672598-0.672598 0.927187-1.599785 1.599785-2.303346 0.512619-0.479935 1.056202-0.832576 1.567101-1.343475C757.759656 538.879828 759.199462 535.391265 761.055557 532.128047z"
              fill="#ffffff"
              p-id="3974"
            />
          </svg>
        </div>
        <div @mouseover="is_hover_category = true">
          <p>耳机 音响</p>
          <svg
            t="1694692637616"
            class="icon"
            viewBox="0 0 1024 1024"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            p-id="3973"
            width="200"
            height="200"
          >
            <path
              d="M761.055557 532.128047c0.512619-0.992555 1.343475-1.823411 1.792447-2.848649 8.800538-18.304636 5.919204-40.703346-9.664077-55.424808L399.935923 139.743798c-19.264507-18.208305-49.631179-17.344765-67.872168 1.888778-18.208305 19.264507-17.375729 49.631179 1.888778 67.872168l316.960409 299.839269L335.199677 813.631716c-19.071845 18.399247-19.648112 48.767639-1.247144 67.872168 9.407768 9.791372 21.984142 14.688778 34.560516 14.688778 12.000108 0 24.000215-4.479398 33.311652-13.439914l350.048434-337.375729c0.672598-0.672598 0.927187-1.599785 1.599785-2.303346 0.512619-0.479935 1.056202-0.832576 1.567101-1.343475C757.759656 538.879828 759.199462 535.391265 761.055557 532.128047z"
              fill="#ffffff"
              p-id="3974"
            />
          </svg>
        </div>
        <div @mouseover="is_hover_category =true">
          <p>健康 儿童</p>
          <svg
            t="1694692637616"
            class="icon"
            viewBox="0 0 1024 1024"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            p-id="3973"
            width="200"
            height="200"
          >
            <path
              d="M761.055557 532.128047c0.512619-0.992555 1.343475-1.823411 1.792447-2.848649 8.800538-18.304636 5.919204-40.703346-9.664077-55.424808L399.935923 139.743798c-19.264507-18.208305-49.631179-17.344765-67.872168 1.888778-18.208305 19.264507-17.375729 49.631179 1.888778 67.872168l316.960409 299.839269L335.199677 813.631716c-19.071845 18.399247-19.648112 48.767639-1.247144 67.872168 9.407768 9.791372 21.984142 14.688778 34.560516 14.688778 12.000108 0 24.000215-4.479398 33.311652-13.439914l350.048434-337.375729c0.672598-0.672598 0.927187-1.599785 1.599785-2.303346 0.512619-0.479935 1.056202-0.832576 1.567101-1.343475C757.759656 538.879828 759.199462 535.391265 761.055557 532.128047z"
              fill="#ffffff"
              p-id="3974"
            />
          </svg>
        </div>
        <div @mouseover="is_hover_category = true">
          <p>生活 箱子包</p>
          <svg
            t="1694692637616"
            class="icon"
            viewBox="0 0 1024 1024"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            p-id="3973"
            width="200"
            height="200"
          >
            <path
              d="M761.055557 532.128047c0.512619-0.992555 1.343475-1.823411 1.792447-2.848649 8.800538-18.304636 5.919204-40.703346-9.664077-55.424808L399.935923 139.743798c-19.264507-18.208305-49.631179-17.344765-67.872168 1.888778-18.208305 19.264507-17.375729 49.631179 1.888778 67.872168l316.960409 299.839269L335.199677 813.631716c-19.071845 18.399247-19.648112 48.767639-1.247144 67.872168 9.407768 9.791372 21.984142 14.688778 34.560516 14.688778 12.000108 0 24.000215-4.479398 33.311652-13.439914l350.048434-337.375729c0.672598-0.672598 0.927187-1.599785 1.599785-2.303346 0.512619-0.479935 1.056202-0.832576 1.567101-1.343475C757.759656 538.879828 759.199462 535.391265 761.055557 532.128047z"
              fill="#ffffff"
              p-id="3974"
            />
          </svg>
        </div>
        <div @mouseover="is_hover_category = true">
          <p>智能 路由器</p>
          <svg
            t="1694692637616"
            class="icon"
            viewBox="0 0 1024 1024"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            p-id="3973"
            width="200"
            height="200"
          >
            <path
              d="M761.055557 532.128047c0.512619-0.992555 1.343475-1.823411 1.792447-2.848649 8.800538-18.304636 5.919204-40.703346-9.664077-55.424808L399.935923 139.743798c-19.264507-18.208305-49.631179-17.344765-67.872168 1.888778-18.208305 19.264507-17.375729 49.631179 1.888778 67.872168l316.960409 299.839269L335.199677 813.631716c-19.071845 18.399247-19.648112 48.767639-1.247144 67.872168 9.407768 9.791372 21.984142 14.688778 34.560516 14.688778 12.000108 0 24.000215-4.479398 33.311652-13.439914l350.048434-337.375729c0.672598-0.672598 0.927187-1.599785 1.599785-2.303346 0.512619-0.479935 1.056202-0.832576 1.567101-1.343475C757.759656 538.879828 759.199462 535.391265 761.055557 532.128047z"
              fill="#ffffff"
              p-id="3974"
            />
          </svg>
        </div>
        <div @mouseover="is_hover_category = true">
          <p>电源 配件</p>
          <svg
            t="1694692637616"
            class="icon"
            viewBox="0 0 1024 1024"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            p-id="3973"
            width="200"
            height="200"
          >
            <path
              d="M761.055557 532.128047c0.512619-0.992555 1.343475-1.823411 1.792447-2.848649 8.800538-18.304636 5.919204-40.703346-9.664077-55.424808L399.935923 139.743798c-19.264507-18.208305-49.631179-17.344765-67.872168 1.888778-18.208305 19.264507-17.375729 49.631179 1.888778 67.872168l316.960409 299.839269L335.199677 813.631716c-19.071845 18.399247-19.648112 48.767639-1.247144 67.872168 9.407768 9.791372 21.984142 14.688778 34.560516 14.688778 12.000108 0 24.000215-4.479398 33.311652-13.439914l350.048434-337.375729c0.672598-0.672598 0.927187-1.599785 1.599785-2.303346 0.512619-0.479935 1.056202-0.832576 1.567101-1.343475C757.759656 538.879828 759.199462 535.391265 761.055557 532.128047z"
              fill="#ffffff"
              p-id="3974"
            />
          </svg>
        </div>
      </div>
      <div v-if="is_hover_category" class="hover_category" @mouseleave="is_hover_category = false">
        <div>时装</div>
        <div>时装</div>
        <div>时装</div>
        <div>时装</div>
        <div>时装</div>
        <div>时装</div>
        <div>时装</div>
        <div>时装</div>
        <div>时装</div>
      </div>

    </div>
    <!--图片信息-->
    <div style="width: 1220px;margin: 10px auto" class="image-info">
      <div class="info">
        <div>
          <span />
          <span />
        </div>
      </div>
      <div class="image">
        <template v-for="item in carouselItemList">
          <el-image style="width: 318px;height: 170px;" :src="item.url" />
        </template>
      </div>
    </div>

    <!--详细商品-->
    <div class="product-detail">
      <img
        src="https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/4d1e0c3e71628cf1df6c941d7a908637.jpeg?thumb=1&w=1839&h=180&f=webp&q=90"
        alt=""
      >

      <div class="product-info">
        <div class="title">
          <p>手机</p>
          <span @click="toProductList(1)">查看更多 >></span>
        </div>
        <div class="detail">
          <img
            src="https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/9613bd4127fede053f8ba5049eff0392.jpeg?thumb=1&w=351&h=921&f=webp&q=90"
          >
          <div class="product_list">
            <div class="product">
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202308111609_92d3c7105793aed5afd3e609132840c9.jpg?thumb=1&w=300&h=300&f=webp&q=90"
              >
              <p class="name">Xiaomi MIX Fold 3</p>
              <p class="desc">轻薄折叠屏丨徕卡光学丨全焦段四摄</p>
              <span> 2399 元起 </span>
            </div>
            <div class="product">
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202308111609_92d3c7105793aed5afd3e609132840c9.jpg?thumb=1&w=300&h=300&f=webp&q=90"
              >
              <p class="name">Xiaomi MIX Fold 3</p>
              <p class="desc">轻薄折叠屏丨徕卡光学丨全焦段四摄</p>
              <span> 2399 元起 </span>
            </div>
            <div class="product">
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202308111609_92d3c7105793aed5afd3e609132840c9.jpg?thumb=1&w=300&h=300&f=webp&q=90"
              >
              <p class="name">Xiaomi MIX Fold 3</p>
              <p class="desc">轻薄折叠屏丨徕卡光学丨全焦段四摄</p>
              <span> 2399 元起 </span>
            </div>
            <div class="product">
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202308111609_92d3c7105793aed5afd3e609132840c9.jpg?thumb=1&w=300&h=300&f=webp&q=90"
              >
              <p class="name">Xiaomi MIX Fold 3</p>
              <p class="desc">轻薄折叠屏丨徕卡光学丨全焦段四摄</p>
              <span> 2399 元起 </span>
            </div>
            <div class="product">
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202308111609_92d3c7105793aed5afd3e609132840c9.jpg?thumb=1&w=300&h=300&f=webp&q=90"
              >
              <p class="name">Xiaomi MIX Fold 3</p>
              <p class="desc">轻薄折叠屏丨徕卡光学丨全焦段四摄</p>
              <span> 2399 元起 </span>
            </div>
            <div class="product">
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202308111609_92d3c7105793aed5afd3e609132840c9.jpg?thumb=1&w=300&h=300&f=webp&q=90"
              >
              <p class="name">Xiaomi MIX Fold 3</p>
              <p class="desc">轻薄折叠屏丨徕卡光学丨全焦段四摄</p>
              <span> 2399 元起 </span>
            </div>
            <div class="product">
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202308111609_92d3c7105793aed5afd3e609132840c9.jpg?thumb=1&w=300&h=300&f=webp&q=90"
              >
              <p class="name">Xiaomi MIX Fold 3</p>
              <p class="desc">轻薄折叠屏丨徕卡光学丨全焦段四摄</p>
              <span> 2399 元起 </span>
            </div>
            <div class="product">
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202308111609_92d3c7105793aed5afd3e609132840c9.jpg?thumb=1&w=300&h=300&f=webp&q=90"
              >
              <p class="name">Xiaomi MIX Fold 3</p>
              <p class="desc">轻薄折叠屏丨徕卡光学丨全焦段四摄</p>
              <span> 2399 元起 </span>
            </div>
          </div>
        </div>
      </div>

      <div class="product-info">
        <div class="title">
          <p>手机</p>
          <span @click="toProductList(1)">查看更多 >></span>
        </div>
        <div class="detail">
          <img
            src="https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/9613bd4127fede053f8ba5049eff0392.jpeg?thumb=1&w=351&h=921&f=webp&q=90"
          >
          <div class="product_list">
            <div class="product">
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202308111609_92d3c7105793aed5afd3e609132840c9.jpg?thumb=1&w=300&h=300&f=webp&q=90"
              >
              <p class="name">Xiaomi MIX Fold 3</p>
              <p class="desc">轻薄折叠屏丨徕卡光学丨全焦段四摄</p>
              <span> 2399 元起 </span>
            </div>
            <div class="product">
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202308111609_92d3c7105793aed5afd3e609132840c9.jpg?thumb=1&w=300&h=300&f=webp&q=90"
              >
              <p class="name">Xiaomi MIX Fold 3</p>
              <p class="desc">轻薄折叠屏丨徕卡光学丨全焦段四摄</p>
              <span> 2399 元起 </span>
            </div>
            <div class="product">
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202308111609_92d3c7105793aed5afd3e609132840c9.jpg?thumb=1&w=300&h=300&f=webp&q=90"
              >
              <p class="name">Xiaomi MIX Fold 3</p>
              <p class="desc">轻薄折叠屏丨徕卡光学丨全焦段四摄</p>
              <span> 2399 元起 </span>
            </div>
            <div class="product">
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202308111609_92d3c7105793aed5afd3e609132840c9.jpg?thumb=1&w=300&h=300&f=webp&q=90"
              >
              <p class="name">Xiaomi MIX Fold 3</p>
              <p class="desc">轻薄折叠屏丨徕卡光学丨全焦段四摄</p>
              <span> 2399 元起 </span>
            </div>
            <div class="product">
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202308111609_92d3c7105793aed5afd3e609132840c9.jpg?thumb=1&w=300&h=300&f=webp&q=90"
              >
              <p class="name">Xiaomi MIX Fold 3</p>
              <p class="desc">轻薄折叠屏丨徕卡光学丨全焦段四摄</p>
              <span> 2399 元起 </span>
            </div>
            <div class="product">
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202308111609_92d3c7105793aed5afd3e609132840c9.jpg?thumb=1&w=300&h=300&f=webp&q=90"
              >
              <p class="name">Xiaomi MIX Fold 3</p>
              <p class="desc">轻薄折叠屏丨徕卡光学丨全焦段四摄</p>
              <span> 2399 元起 </span>
            </div>
            <div class="product">
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202308111609_92d3c7105793aed5afd3e609132840c9.jpg?thumb=1&w=300&h=300&f=webp&q=90"
              >
              <p class="name">Xiaomi MIX Fold 3</p>
              <p class="desc">轻薄折叠屏丨徕卡光学丨全焦段四摄</p>
              <span> 2399 元起 </span>
            </div>
            <div class="product">
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202308111609_92d3c7105793aed5afd3e609132840c9.jpg?thumb=1&w=300&h=300&f=webp&q=90"
              >
              <p class="name">Xiaomi MIX Fold 3</p>
              <p class="desc">轻薄折叠屏丨徕卡光学丨全焦段四摄</p>
              <span> 2399 元起 </span>
            </div>
          </div>
        </div>
      </div>
      <div class="product-info">
        <div class="title">
          <p>手机</p>
          <span @click="toProductList(1)">查看更多 >></span>
        </div>
        <div class="detail">
          <img
            src="https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/9613bd4127fede053f8ba5049eff0392.jpeg?thumb=1&w=351&h=921&f=webp&q=90"
          >
          <div class="product_list">
            <div class="product">
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202308111609_92d3c7105793aed5afd3e609132840c9.jpg?thumb=1&w=300&h=300&f=webp&q=90"
              >
              <p class="name">Xiaomi MIX Fold 3</p>
              <p class="desc">轻薄折叠屏丨徕卡光学丨全焦段四摄</p>
              <span> 2399 元起 </span>
            </div>
            <div class="product">
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202308111609_92d3c7105793aed5afd3e609132840c9.jpg?thumb=1&w=300&h=300&f=webp&q=90"
              >
              <p class="name">Xiaomi MIX Fold 3</p>
              <p class="desc">轻薄折叠屏丨徕卡光学丨全焦段四摄</p>
              <span> 2399 元起 </span>
            </div>
            <div class="product">
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202308111609_92d3c7105793aed5afd3e609132840c9.jpg?thumb=1&w=300&h=300&f=webp&q=90"
              >
              <p class="name">Xiaomi MIX Fold 3</p>
              <p class="desc">轻薄折叠屏丨徕卡光学丨全焦段四摄</p>
              <span> 2399 元起 </span>
            </div>
            <div class="product">
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202308111609_92d3c7105793aed5afd3e609132840c9.jpg?thumb=1&w=300&h=300&f=webp&q=90"
              >
              <p class="name">Xiaomi MIX Fold 3</p>
              <p class="desc">轻薄折叠屏丨徕卡光学丨全焦段四摄</p>
              <span> 2399 元起 </span>
            </div>
            <div class="product">
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202308111609_92d3c7105793aed5afd3e609132840c9.jpg?thumb=1&w=300&h=300&f=webp&q=90"
              >
              <p class="name">Xiaomi MIX Fold 3</p>
              <p class="desc">轻薄折叠屏丨徕卡光学丨全焦段四摄</p>
              <span> 2399 元起 </span>
            </div>
            <div class="product">
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202308111609_92d3c7105793aed5afd3e609132840c9.jpg?thumb=1&w=300&h=300&f=webp&q=90"
              >
              <p class="name">Xiaomi MIX Fold 3</p>
              <p class="desc">轻薄折叠屏丨徕卡光学丨全焦段四摄</p>
              <span> 2399 元起 </span>
            </div>
            <div class="product">
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202308111609_92d3c7105793aed5afd3e609132840c9.jpg?thumb=1&w=300&h=300&f=webp&q=90"
              >
              <p class="name">Xiaomi MIX Fold 3</p>
              <p class="desc">轻薄折叠屏丨徕卡光学丨全焦段四摄</p>
              <span> 2399 元起 </span>
            </div>
            <div class="product">
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202308111609_92d3c7105793aed5afd3e609132840c9.jpg?thumb=1&w=300&h=300&f=webp&q=90"
              >
              <p class="name">Xiaomi MIX Fold 3</p>
              <p class="desc">轻薄折叠屏丨徕卡光学丨全焦段四摄</p>
              <span> 2399 元起 </span>
            </div>
          </div>
        </div>
      </div>

      <img
        src="https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/4d1e0c3e71628cf1df6c941d7a908637.jpeg?thumb=1&w=1839&h=180&f=webp&q=90"
        alt=""
      >
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const carouselItemList = ref([
  { url: 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/7bfda2a88bc0ce1f17b44cce285be2d4.jpg?thumb=1&w=1839&h=690&f=webp&q=90' },
  { url: 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/da7c749aa3ce4f675ad44d7ca80ad749.jpg?thumb=1&w=1839&h=690&f=webp&q=90' },
  { url: 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/309d66406483fd2da682e014cc772ef3.jpg?thumb=1&w=1839&h=690&f=webp&q=90' }
])

const is_hover_category = ref(false)

const toProductList = (id) => {
  router.push({ name: 'ProductList' })
}
</script>

<style scoped lang="scss">
.mall_carousel_category {
  height: 460px;
  width: 100%;
  position: relative;
}

:deep(.el-carousel__container ) {
  height: 100%;
}

.mall_category {
  height: 460px;
  width: 234px;
  background-color: #9f9d9d;
  position: absolute;
  top: 0;
  left: 0;
  z-index: 999;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  padding: 30px 0;
  box-sizing: border-box;

  & > div {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 20px;
    box-sizing: border-box;
    width: 100%;
    height: 40px;

    p {
      font-size: 14px;
      color: white;
      cursor: pointer;
    }

    & > svg {
      width: 16px;
      height: 16px;
    }
  }

  & > div:hover {
    background-color: #ff6700;
  }
}

.mall_carousel_category > .hover_category {
  box-shadow: 0 8px 16px rgba(0, 0, 0, .18);
  height: 460px;
  width: 234px;
  background-color: white;
  display: inline-block;
  position: relative;
  padding: 20px 0;
  box-sizing: border-box;
  top: -460px;
  left: 234px;

  & > div {
    width: 234px;
    height: 40px;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 14px;
    letter-spacing: 2px;
    float: left;
  }

  & > div:hover {
    color: #ff6700;
  }
}

.image-info {
  display: flex;
  margin-bottom: 20px !important;
  justify-content: space-between;

  .info {
    width: 234px;
    height: 170px;
    background-color: blue;
  }

  .image {
    width: 976px;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }
}

.product-detail {
  width: 100vw;
  margin: 0 auto;
  background-color: #f5f5f5;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;

  & > img {
    width: 1220px;
    margin: 20px 0 40px 0;
    height: 140px;
  }

  > .product-info {
    width: 1220px;
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-bottom: 20px;

    .title {
      width: 100%;
      display: flex;
      flex-direction: row;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 15px;

      p {
        font-size: 22px;
        font-weight: 300;
        color: #595c59;
      }

      span {
        position: relative;
        top: 6px;
        color: #595c59;
        transition: all 0.3s;
      }

      span:hover {
        color: #ff6700;
        cursor: pointer;
      }
    }

    .detail {
      width: 100%;
      height: 610px;
      display: flex;
      justify-content: space-between;

      img {
        width: 234px;
        height: 100%
      }

      .product_list {
        width: 976px;
        display: flex;
        justify-content: space-between;
        flex-wrap: wrap;
        align-content: space-between;

        .product {
          background-color: white;
          width: 235px;
          height: 300px;
          display: flex;
          align-items: center;
          flex-direction: column;
          justify-content: space-between;
          transition: all 0.3s;

          padding-bottom: 40px;
          box-sizing: border-box;

          img {
            width: 160px;
            height: 160px;
          }

          .name {
            font-size: 14px;
          }

          .desc {
            font-size: 12px;
            color: #9f9d9d;
            position: relative;
            top: -8px;
          }

          span {
            color: #ff6700;
            font-size: 14px;
          }

        }

        .product:hover {
          cursor: pointer;
          box-shadow: 0px 0px 37px -7px rgba(0, 0, 0, 0.18);
        }
      }
    }
  }
}
</style>
