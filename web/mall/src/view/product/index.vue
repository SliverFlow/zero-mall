<template>
  <div class="content_container">
    <div class="title">
      <p>首页</p>
      <span> > </span>
      <p>全部商品</p>
    </div>
    <!--分类显示-->
    <div class="category_container">
      <div class="category_content">
        <el-collapse accordion>
          <el-collapse-item name="active">
            <template #title>
              <div class="category_item">
                <p class="main">主分类</p>
                <a href="" class="secondary">次分类</a>
                <a href="" class="secondary">次分类</a>
                <a href="" class="secondary">次分类</a>
                <a href="" class="secondary">次分类</a>
                <a href="" class="secondary">次分类</a>
                <a href="" class="secondary">次分类</a>
                <a href="" class="secondary">次分类</a>
                <a href="" class="secondary">次分类</a>
              </div>
            </template>
            <div class="category_item" style="margin-bottom: 10px">
              <p>主分类</p>
              <a href="">次分类</a>
              <a href="">次分类</a>
              <a href="">次分类</a>
              <a href="">次分类</a>
              <a href="">次分类</a>
              <a href="">次分类</a>
              <a href="">次分类</a>
              <a href="">次分类</a>
              <a href="">次分类</a>
            </div>
            <div class="category_item" style="margin-bottom: 10px">
              <p>主分类</p>
              <a href="">次分类</a>
              <a href="">次分类</a>
              <a href="">次分类</a>
              <a href="">次分类</a>
              <a href="">次分类</a>
              <a href="">次分类</a>
              <a href="">次分类</a>
              <a href="">次分类</a>
              <a href="">次分类</a>
            </div>
            <div class="category_item" style="margin-bottom: 10px">
              <p>主分类</p>
              <a href="">次分类</a>
              <a href="">次分类</a>
              <a href="">次分类</a>
              <a href="">次分类</a>
              <a href="">次分类</a>
              <a href="">次分类</a>
              <a href="">次分类</a>
              <a href="">次分类</a>
              <a href="">次分类</a>
            </div>
          </el-collapse-item>
        </el-collapse>
      </div>
    </div>
    <!--筛选条件-->
    <div class="screen">
      <div class="content">
        <p>综合</p>
        <span>|</span>
        <p>销量</p>
        <span>|</span>
        <p>价格</p>
      </div>
    </div>
    <!--商品列表-->
    <div class="product_list">
      <div class="content">
        <div class="prod" @click="toProductDetail(1)">
          <img
            src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202212242022_f9456bf36437f2888d3a7b733319743f.png"
            alt=""
          >
          <el-text style="width: 240px;" truncated>两挂一柜（3匹）全屋空调套餐 新风系列</el-text>
          <span class="price">13997元</span>
          <ul>
            <li>
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202303281923_3b9ebd24f96ac04d2b2bc9addfb78051.png"
                alt=""
              >
            </li>
          </ul>
          <div class="tag">
            <span>分期</span>
          </div>
        </div>
        <div class="prod" @click="toProductDetail(1)">
          <img
            src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202212242022_f9456bf36437f2888d3a7b733319743f.png"
            alt=""
          >
          <el-text style="width: 240px;" truncated>两挂一柜（3匹）全屋空调套餐 新风系列</el-text>
          <span class="price">13997元</span>
          <ul>
            <li>
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202303281923_3b9ebd24f96ac04d2b2bc9addfb78051.png"
                alt=""
              >
            </li>
          </ul>
          <div class="tag">
            <span>分期</span>
          </div>
        </div>
        <div class="prod" @click="toProductDetail(1)">
          <img
            src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202212242022_f9456bf36437f2888d3a7b733319743f.png"
            alt=""
          >
          <el-text style="width: 240px;" truncated>两挂一柜（3匹）全屋空调套餐 新风系列</el-text>
          <span class="price">13997元</span>
          <ul>
            <li>
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202303281923_3b9ebd24f96ac04d2b2bc9addfb78051.png"
                alt=""
              >
            </li>
          </ul>
          <div class="tag">
            <span>分期</span>
          </div>
        </div>
        <div class="prod" @click="toProductDetail(1)">
          <img
            src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202212242022_f9456bf36437f2888d3a7b733319743f.png"
            alt=""
          >
          <el-text style="width: 240px;" truncated>两挂一柜（3匹）全屋空调套餐 新风系列</el-text>
          <span class="price">13997元</span>
          <ul>
            <li>
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202303281923_3b9ebd24f96ac04d2b2bc9addfb78051.png"
                alt=""
              >
            </li>
          </ul>
          <div class="tag">
            <span>分期</span>
          </div>
        </div>
        <div class="prod">
          <img
            src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202212242022_f9456bf36437f2888d3a7b733319743f.png"
            alt=""
          >
          <el-text style="width: 240px;" truncated>两挂一柜（3匹）全屋空调套餐 新风系列</el-text>
          <span class="price">13997元</span>
          <ul>
            <li>
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202303281923_3b9ebd24f96ac04d2b2bc9addfb78051.png"
                alt=""
              >
            </li>
          </ul>
          <div class="tag">
            <span>分期</span>
          </div>
        </div>
        <div class="prod">
          <img
            src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202212242022_f9456bf36437f2888d3a7b733319743f.png"
            alt=""
          >
          <el-text style="width: 240px;" truncated>两挂一柜（3匹）全屋空调套餐 新风系列</el-text>
          <span class="price">13997元</span>
          <ul>
            <li>
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202303281923_3b9ebd24f96ac04d2b2bc9addfb78051.png"
                alt=""
              >
            </li>
          </ul>
          <div class="tag">
            <span>分期</span>
          </div>
        </div>
        <div class="prod">
          <img
            src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202212242022_f9456bf36437f2888d3a7b733319743f.png"
            alt=""
          >
          <el-text style="width: 240px;" truncated>两挂一柜（3匹）全屋空调套餐 新风系列</el-text>
          <span class="price">13997元</span>
          <ul>
            <li>
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202303281923_3b9ebd24f96ac04d2b2bc9addfb78051.png"
                alt=""
              >
            </li>
          </ul>
          <div class="tag">
            <span>分期</span>
          </div>
        </div>
        <div class="prod">
          <img
            src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202212242022_f9456bf36437f2888d3a7b733319743f.png"
            alt=""
          >
          <el-text style="width: 240px;" truncated>两挂一柜（3匹）全屋空调套餐 新风系列</el-text>
          <span class="price">13997元</span>
          <ul>
            <li>
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202303281923_3b9ebd24f96ac04d2b2bc9addfb78051.png"
                alt=""
              >
            </li>
          </ul>
          <div class="tag">
            <span>分期</span>
          </div>
        </div>
        <div class="prod">
          <img
            src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202212242022_f9456bf36437f2888d3a7b733319743f.png"
            alt=""
          >
          <el-text style="width: 240px;" truncated>两挂一柜（3匹）全屋空调套餐 新风系列</el-text>
          <span class="price">13997元</span>
          <ul>
            <li>
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202303281923_3b9ebd24f96ac04d2b2bc9addfb78051.png"
                alt=""
              >
            </li>
          </ul>
          <div class="tag">
            <span>分期</span>
          </div>
        </div>
        <div class="prod">
          <img
            src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202212242022_f9456bf36437f2888d3a7b733319743f.png"
            alt=""
          >
          <el-text style="width: 240px;" truncated>两挂一柜（3匹）全屋空调套餐 新风系列</el-text>
          <span class="price">13997元</span>
          <ul>
            <li>
              <img
                src="https://cdn.cnbj1.fds.api.mi-img.com/nr-pub/202303281923_3b9ebd24f96ac04d2b2bc9addfb78051.png"
                alt=""
              >
            </li>
          </ul>
          <div class="tag">
            <span>分期</span>
          </div>
        </div>
      </div>
    </div>
    <!--分页-->
    <div class="pagination">
      <el-pagination
        class="page_con"
        layout="total, sizes, prev, pager, next, jumper"
        :current-page="page"
        :page-size="pageSize"
        :page-sizes="[10, 30, 50, 100]"
        :total="total"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
      />
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const pageSize = ref(20)
const page = ref(10)
const total = ref(1000)

const handleCurrentChange = (val) => {}
const handleSizeChange = (val) => {}

const toProductDetail = (id) => {
  router.push({ name: 'ProductDetail' })
}

</script>

<style scoped lang="scss">

.content_container {
  width: 100%;
  background-color: #f5f5f5;
  display: flex;
  flex-direction: column;
  padding-bottom: 20px;

  .title {
    width: 1220px;
    height: 40px;
    margin: 0 auto;
    display: flex;
    align-items: center;
    color: #757575;

    p {
      font-size: 12px;
    }

    span {
      font-size: 12px;
      margin: 0 8px;
    }
  }

  .category_container {
    width: 100%;
    background-color: white;

    .category_content {
      width: 1220px;
      margin: 0 auto;

      .category_item {
        display: flex;
        align-items: center;

        p {
          font-size: 15px;
          margin-right: 10px;
        }

        a {
          font-size: 14px;
          margin-left: 14px;
          transition: all 0.3s;
        }

        a:hover {
          color: #ff6700;
        }
      }
    }
  }

  .screen {
    width: 100%;
    height: 60px;
    background-color: #f5f5f5;

    .content {
      width: 1220px;
      margin: 10px auto;
      height: 100%;
      display: flex;
      align-items: center;

      p {
        font-size: 14px;
        color: #424242;
        transition: all 0.3s;
        cursor: pointer;
      }

      p:hover {
        color: #ff6700;
      }

      span {
        margin: 0 35px;
        font-size: 12px;
        color: rgba(159, 157, 157, 0.65);
      }
    }
  }

  .product_list {
    width: 100%;
    background-color: #f5f5f5;

    .content {
      width: 1220px;
      background-color: bisque;
      margin: 0 auto;

      .prod {
        width: 297px;
        height: 430px;
        background-color: white;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        cursor: pointer;
        margin-right: 10px;
        margin-bottom: 10px;
        float: left;
        transition: all 0.3s;

        img {
          width: 200px;
          height: 200px;
          margin-bottom: 20px;
        }

        .price {
          font-size: 14px;
          color: #ff6700;
          margin-bottom: 14px;
        }

        .tag {
          display: flex;
          align-items: center;
          justify-content: center;

          span {
            height: 18px;
            padding: 2px 6px;
            font-size: 13px;
            line-height: 18px;
            background: radial-gradient(circle, #51a4eb, #62b9f1);
            color: white;
            border-radius: 2px;
          }
        }

        ul > li > img {
          width: 34px;
          height: 34px;
          border: 1px solid #ccc;
          transition: all 0.2s;
          margin-bottom: 14px;

          &:hover {
            border: 1px solid #ff6700;
          }
        }
      }

      .prod:hover {
        cursor: pointer;
        box-shadow: 0px 0px 37px -7px rgba(0, 0, 0, 0.18);
      }

      .prod:nth-child(4n) {
        margin-right: 0;
      }
    }
  }

  .pagination {
    width: 100%;
    background-color: #f5f5f5;
    display: flex;
    align-items: center;
    justify-content: center;
    margin: 18px 0 40px 0;
  }
}

:deep(.el-collapse-item__content) {
  padding-bottom: 8px;
}

:deep(.el-pager li.is-active) {
  color: #ff6700;
}

:deep(.el-select .el-input.is-focus .el-input__wrapper) {
  box-shadow: 0 0 0 1px #ff6700 inset !important;
}

:deep(.el-input__wrapper) {
  background-color: #f5f5f5;
}
:deep(.el-input__wrapper.is-focus) {
  box-shadow: 0 0 0 1px #ff6700 inset !important;
}

:deep(.el-pager li) {
  background-color: #f5f5f5;
  margin-left: 0.25rem;
  margin-right: 0.25rem;
  border-radius: 0.25rem;
  border-width: 1px;
  border-style: solid;
  --tw-border-opacity: 1;
  border-color: rgb(209 213 219 / var(--tw-border-opacity));
  font-size: .875rem;
  line-height: 1.25rem;
  --tw-text-opacity: 1;
  color: rgb(75 85 99 / var(--tw-text-opacity));
  transition: all 0.2s;
}
:deep(.el-pager li):hover {
  color: #ff6700;
}

:deep(.btn-prev) , :deep(.btn-next){
  background-color: #f5f5f5 !important;
  margin-left: 0.25rem;
  margin-right: 0.25rem;
  border-radius: 0.25rem;
  border-width: 1px;
  border-style: solid;
  --tw-border-opacity: 1;
  border-color: rgb(209 213 219 / var(--tw-border-opacity));
  font-size: .875rem;
  line-height: 1.25rem;
  --tw-text-opacity: 1;
  color: rgb(75 85 99 / var(--tw-text-opacity));
  transition: all 0.2s;
}

:deep(.btn-prev):hover,:deep(.btn-next):hover {
  border-color: #ff6700;
  color: #ff6700;
}

</style>
