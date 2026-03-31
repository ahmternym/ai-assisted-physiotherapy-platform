<template>
  <!--
    @click="onCardClick" kullanıyoruz — $emit('click') Vue 3'te native DOM
    click olayıyla çakışır ve parent'taki @click handler'ı çift tetiklenir
    ya da id değeri yerine MouseEvent objesi iletilir. 'select-physio' özel
    event adı bu sorunu tamamen ortadan kaldırır.
  -->
  <div class="physio-card" @click="onCardClick">
    <div class="card-avatar" :style="{ background: avatarGradient }">
      {{ initials }}
    </div>
    <div class="card-info">
      <h3 class="card-name">{{ physio.name }}</h3>
      <div class="card-meta">
        <span class="meta-tag location">
          <span class="meta-icon">📍</span>
          {{ physio.district }}, {{ physio.city }}
        </span>
        <span class="meta-tag phone">
          <span class="meta-icon">📞</span>
          {{ physio.phone }}
        </span>
      </div>
      <div v-if="physio.specialties && physio.specialties.length" class="card-specialties">
        <span
          v-for="(s, i) in physio.specialties.slice(0, 2)"
          :key="i"
          class="spec-chip"
        >{{ s }}</span>
        <span v-if="physio.specialties.length > 2" class="spec-more">
          +{{ physio.specialties.length - 2 }}
        </span>
      </div>
    </div>
    <div class="card-cta">
      <span class="cta-text">Detay</span>
      <span class="cta-arrow">›</span>
    </div>
  </div>
</template>

<script>
const GRADIENTS = [
  'linear-gradient(135deg,#3b82f6,#1d4ed8)',
  'linear-gradient(135deg,#8b5cf6,#6d28d9)',
  'linear-gradient(135deg,#ec4899,#be185d)',
  'linear-gradient(135deg,#10b981,#047857)',
  'linear-gradient(135deg,#f59e0b,#b45309)',
  'linear-gradient(135deg,#06b6d4,#0e7490)',
]

export default {
  name: 'PhysiotherapistCard',

  props: {
    physio: { type: Object, required: true }
  },

  // Vue 3: emits seçeneği ile özel olayları açıkça tanımla
  emits: ['select-physio'],

  computed: {
    initials() {
      return this.physio.name
        .split(' ')
        .map(n => n[0])
        .join('')
        .substring(0, 2)
        .toUpperCase()
    },
    avatarGradient() {
      return GRADIENTS[this.physio.id % GRADIENTS.length]
    }
  },

  methods: {
    onCardClick() {
      // Fizyoterapist ID'sini parent'a ilet
      this.$emit('select-physio', this.physio.id)
    }
  }
}
</script>

<style scoped>
.physio-card {
  display: flex;
  align-items: center;
  gap: 1rem;
  background: #fff;
  border: 1.5px solid #e8eef6;
  border-radius: 16px;
  padding: 1rem 1.2rem;
  cursor: pointer;
  transition: all 0.22s cubic-bezier(.4,0,.2,1);
  box-shadow: 0 2px 8px rgba(30,58,95,0.06);
  position: relative;
  overflow: hidden;
}

.physio-card::before {
  content: '';
  position: absolute;
  left: 0; top: 0; bottom: 0;
  width: 4px;
  background: linear-gradient(180deg, #3b82f6, #8b5cf6);
  border-radius: 4px 0 0 4px;
  opacity: 0;
  transition: opacity 0.22s;
}

.physio-card:hover {
  border-color: #bfdbfe;
  box-shadow: 0 8px 28px rgba(59,130,246,0.14);
  transform: translateY(-3px);
}

.physio-card:hover::before { opacity: 1; }

.card-avatar {
  width: 54px;
  height: 54px;
  border-radius: 50%;
  color: white;
  font-weight: 800;
  font-size: 1.1rem;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  box-shadow: 0 3px 12px rgba(0,0,0,0.15);
  letter-spacing: 0.05em;
}

.card-info { flex: 1; min-width: 0; }

.card-name {
  margin: 0 0 0.35rem;
  font-size: 1rem;
  color: #0f172a;
  font-weight: 700;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.card-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
  margin-bottom: 0.5rem;
}

.meta-tag {
  display: inline-flex;
  align-items: center;
  gap: 0.25rem;
  font-size: 0.78rem;
  color: #64748b;
  background: #f8fafc;
  border: 1px solid #e2e8f0;
  border-radius: 99px;
  padding: 0.2rem 0.6rem;
}
.meta-icon { font-size: 0.75rem; }

.card-specialties { display: flex; flex-wrap: wrap; gap: 0.35rem; }

.spec-chip {
  font-size: 0.72rem;
  font-weight: 600;
  background: #eff6ff;
  color: #2563eb;
  border-radius: 6px;
  padding: 0.15rem 0.5rem;
}

.spec-more {
  font-size: 0.72rem;
  font-weight: 600;
  background: #f0fdf4;
  color: #16a34a;
  border-radius: 6px;
  padding: 0.15rem 0.5rem;
}

.card-cta {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.1rem;
  flex-shrink: 0;
}

.cta-text {
  font-size: 0.7rem;
  color: #94a3b8;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.04em;
}

.cta-arrow {
  font-size: 1.6rem;
  color: #cbd5e1;
  line-height: 1;
  transition: color 0.2s, transform 0.2s;
}

.physio-card:hover .cta-arrow {
  color: #3b82f6;
  transform: translateX(3px);
}
</style>
