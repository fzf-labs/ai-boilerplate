<script setup lang="ts">
const trustLogos = ['云岚会员', '拾光零售', '栖岸内容', '星海社群', '飞霓科技', '诺合服务']

const metrics = [
  { value: '1 套', label: '统一运营后台', note: '权限、审计、报表一体化' },
  { value: '3 端', label: '用户触达入口', note: 'App / H5 / 官网同步体验' },
  { value: '5+', label: '核心业务模块', note: '会员 / 商城 / 内容 / 公众号 / AI' },
  { value: '100%', label: '可扩展架构', note: '模块化前后端' },
]

const securityItems = [
  '多角色权限与菜单级控制',
  '操作日志与审计留痕',
  '订单与支付风控校验',
  '公告与消息触达记录',
]

const testimonials = [
  {
    quote: '会员体系和支付流程一周内跑通，增长团队终于可以专注拉新。',
    name: '林语',
    title: '拾光零售，用户增长负责人',
  },
  {
    quote: '公众号粉丝与商城订单打通后，运营动作更精准，复购明显提升。',
    name: '程远',
    title: '云岚会员，运营经理',
  },
  {
    quote: 'AI 辅助生成内容与图像，节省了大量素材制作时间。',
    name: '顾岚',
    title: '栖岸内容，品牌总监',
  },
]
</script>

<template>
  <div class="page">
    <header class="nav" id="top">
      <div class="container nav__inner">
        <a class="brand" href="#top">
          <span class="brand__mark"></span>
          <span class="brand__name">AI Boilerplate</span>
        </a>
        <nav class="nav__links">
          <a href="#platform">平台能力</a>
          <a href="#solutions">业务模块</a>
          <a href="#security">安全合规</a>
          <a href="#stories">客户故事</a>
          <a href="#contact">联系</a>
        </nav>
        <div class="nav__cta">
          <a class="btn btn--ghost" href="#contact">预约演示</a>
          <a class="btn btn--primary" href="#contact">开始试用</a>
        </div>
      </div>
    </header>

    <main>
      <section class="section hero">
        <div class="container hero__grid">
          <div class="hero__content">
            <p class="pill reveal" style="--delay: 0ms">后台 + App + 官网一体化方案</p>
            <h1 class="hero__title reveal" style="--delay: 80ms">
              一套系统打通会员、商城与内容运营。
            </h1>
            <p class="hero__subtitle reveal" style="--delay: 140ms">
              AI Boilerplate 提供管理后台、会员 App 与官网一体化能力，覆盖商品、订单、支付、
              会员权益、内容发布、公众号运营及 AI 辅助，开箱即用且可持续扩展。
            </p>
            <div class="hero__actions reveal" style="--delay: 200ms">
              <a class="btn btn--primary" href="#contact">申请演示</a>
              <a class="btn btn--ghost" href="#platform">查看平台能力</a>
            </div>
            <div class="hero__highlights">
              <div class="highlight-card reveal" style="--delay: 240ms">
                <p class="highlight-card__value">后台 + App + 官网</p>
                <p class="highlight-card__label">三端协同</p>
              </div>
              <div class="highlight-card reveal" style="--delay: 300ms">
                <p class="highlight-card__value">会员 + 商城</p>
                <p class="highlight-card__label">业务闭环</p>
              </div>
              <div class="highlight-card reveal" style="--delay: 360ms">
                <p class="highlight-card__value">内容 + 触达</p>
                <p class="highlight-card__label">运营提效</p>
              </div>
            </div>
          </div>
          <div class="hero__visual reveal" style="--delay: 220ms">
            <div class="device">
              <div class="device__notch"></div>
              <div class="device__screen">
                <div class="screen__header">
                  <span class="screen__time">09:41</span>
                  <span class="screen__status">会员中心</span>
                </div>
                <div class="screen__cards">
                  <div class="screen__card screen__card--primary">
                    <p class="screen__card-label">会员活跃度</p>
                    <p class="screen__card-value">68%</p>
                    <div class="screen__bar">
                      <span style="width: 68%"></span>
                    </div>
                  </div>
                  <div class="screen__card">
                    <p class="screen__card-label">待发货订单</p>
                    <p class="screen__card-value">24</p>
                    <p class="screen__card-note">今日新增 5</p>
                  </div>
                  <div class="screen__card">
                    <p class="screen__card-label">本月续费率</p>
                    <p class="screen__card-value">18.6%</p>
                    <p class="screen__card-note">较上月 +2.1%</p>
                  </div>
                </div>
                <div class="screen__footer">
                  <div class="screen__chip">权益对比</div>
                  <div class="screen__chip">内容推送</div>
                </div>
              </div>
            </div>
            <div class="floating-card">
              <p class="floating-card__title">订单已完成</p>
              <p class="floating-card__meta">2 分钟前</p>
            </div>
          </div>
        </div>
      </section>

      <section class="section trust">
        <div class="container trust__inner">
          <p class="trust__label">正在使用 AI Boilerplate 的团队</p>
          <div class="trust__logos">
            <span v-for="logo in trustLogos" :key="logo">{{ logo }}</span>
          </div>
        </div>
      </section>

      <section class="section platform" id="platform">
        <div class="container">
          <div class="section-heading">
            <p class="eyebrow">平台能力</p>
            <h2>多端协同，运营链路更清晰。</h2>
            <p>从配置到触达，后台、App 与官网体验一致，降低沟通与执行成本。</p>
          </div>
          <div class="feature-grid">
            <div class="feature-card">
              <div class="feature-card__icon">
                <svg viewBox="0 0 24 24" aria-hidden="true">
                  <path
                    d="M12 3l7 4v5c0 5-3.5 9-7 10-3.5-1-7-5-7-10V7l7-4z"
                  />
                </svg>
              </div>
              <h3>统一管理后台</h3>
              <p>商品、订单、支付、会员、激活码与数据看板集中管理。</p>
            </div>
            <div class="feature-card">
              <div class="feature-card__icon">
                <svg viewBox="0 0 24 24" aria-hidden="true">
                  <path d="M4 7h16M4 12h16M4 17h10" />
                  <path d="M7 4h10v16H7z" />
                </svg>
              </div>
              <h3>App 与官网用户端</h3>
              <p>登录、权益、订单、内容一体化，App / H5 / 官网同步体验。</p>
            </div>
            <div class="feature-card">
              <div class="feature-card__icon">
                <svg viewBox="0 0 24 24" aria-hidden="true">
                  <path d="M12 3v18M3 12h18" />
                  <path d="M7 7l10 10M7 17L17 7" />
                </svg>
              </div>
              <h3>内容与公众号运营</h3>
              <p>素材、粉丝、自动回复与消息通知打通，形成运营闭环。</p>
            </div>
          </div>
        </div>
      </section>

      <section class="section bento" id="solutions">
        <div class="container">
          <div class="section-heading">
            <p class="eyebrow">业务模块</p>
            <h2>覆盖增长与变现的关键环节。</h2>
            <p>把会员、商城、内容与触达串起来，运营动作更可控。</p>
          </div>
          <div class="bento-grid">
            <div class="bento-card bento-card--large">
              <div class="tag">会员增长</div>
              <h3>等级与权益体系</h3>
              <p>会员等级、权益配置、到期提醒与续费流程，提升留存。</p>
              <div class="bento-card__metrics">
                <div>
                  <span>3 分钟</span>
                  <p>完成权益配置</p>
                </div>
                <div>
                  <span>2.1x</span>
                  <p>续费转化提升</p>
                </div>
              </div>
            </div>
            <div class="bento-card">
              <div class="tag">商城</div>
              <h3>商品与订单闭环</h3>
              <p>商品、库存、订单、支付记录与激活码统一管理。</p>
            </div>
            <div class="bento-card">
              <div class="tag">内容</div>
              <h3>素材与内容管理</h3>
              <p>图文、内容详情、消息通知一站式配置与分发。</p>
            </div>
            <div class="bento-card bento-card--wide">
              <div class="tag">AI 运营</div>
              <h3>内容与客服助手</h3>
              <p>内置 AI 文案、图像生成与对话支持，提升运营产能。</p>
              <div class="bento-card__chips">
                <span>智能文案</span>
                <span>图像生成</span>
                <span>对话辅助</span>
              </div>
            </div>
          </div>
        </div>
      </section>

      <section class="section metrics" id="metrics">
        <div class="container metrics__grid">
          <div v-for="metric in metrics" :key="metric.label" class="metric-card">
            <p class="metric-card__value">{{ metric.value }}</p>
            <p class="metric-card__label">{{ metric.label }}</p>
            <p class="metric-card__note">{{ metric.note }}</p>
          </div>
        </div>
      </section>

      <section class="section security" id="security">
        <div class="container security__grid">
          <div>
            <p class="eyebrow">安全合规</p>
            <h2>权限、审计与风控贯穿全流程。</h2>
            <p class="section-text">
              角色权限、操作日志与消息通知闭环，让每一次运营动作可追踪、可验证。
            </p>
            <div class="security__card">
              <div class="security__stat">
                <span>权限体系</span>
                <p>角色、菜单、接口多维度控制。</p>
              </div>
              <div class="security__stat">
                <span>操作审计</span>
                <p>登录与操作记录全程留痕。</p>
              </div>
            </div>
          </div>
          <div class="security__list">
            <h3>合规亮点</h3>
            <ul>
              <li v-for="item in securityItems" :key="item">
                <span class="check-icon">
                  <svg viewBox="0 0 24 24" aria-hidden="true">
                    <path d="M5 12l4 4 10-10" />
                  </svg>
                </span>
                {{ item }}
              </li>
            </ul>
            <a class="link" href="#contact">获取安全方案</a>
          </div>
        </div>
      </section>

      <section class="section stories" id="stories">
        <div class="container">
          <div class="section-heading">
            <p class="eyebrow">客户故事</p>
            <h2>增长与运营团队的共同选择。</h2>
            <p>从会员留存到内容触达，AI Boilerplate 帮助团队快速落地。</p>
          </div>
          <div class="quote-grid">
            <figure v-for="story in testimonials" :key="story.name" class="quote-card">
              <blockquote>"{{ story.quote }}"</blockquote>
              <figcaption>
                <span>{{ story.name }}</span>
                <p>{{ story.title }}</p>
              </figcaption>
            </figure>
          </div>
        </div>
      </section>

      <section class="section cta" id="contact">
        <div class="container">
          <div class="cta__panel">
            <div>
              <p class="eyebrow">开始使用</p>
              <h2>想更快上线并跑通业务吗？</h2>
              <p>预约演示，我们会根据你的业务模块提供落地方案。</p>
            </div>
            <div class="cta__actions">
              <a class="btn btn--primary" href="mailto:hello@ai-boilerplate.com">预约演示</a>
              <a class="btn btn--ghost" href="mailto:sales@ai-boilerplate.com">联系销售</a>
            </div>
          </div>
        </div>
      </section>
    </main>

    <footer class="footer">
      <div class="container footer__inner">
        <div>
          <div class="brand footer__brand">
            <span class="brand__mark"></span>
            <span class="brand__name">AI Boilerplate</span>
          </div>
          <p class="footer__note">企业级运营一体化模板：后台 + App + 官网 + API 一站式交付。</p>
        </div>
        <div class="footer__links">
          <div>
            <h4>产品</h4>
            <a href="#platform">平台能力</a>
            <a href="#solutions">业务模块</a>
            <a href="#security">安全合规</a>
          </div>
          <div>
            <h4>公司</h4>
            <a href="#stories">客户</a>
            <a href="#contact">联系</a>
            <a href="#top">新闻中心</a>
          </div>
          <div>
            <h4>联系</h4>
            <p>沪上路 312 号</p>
            <p>上海</p>
            <p>hello@ai-boilerplate.com</p>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>
