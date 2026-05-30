"""
Statistical comparison infographic: LIU Huanhua vs DJURAEV Akbar
With Sinclair scores — shows how bodyweight adjustment changes the picture.
Output: exactly 1080×1080 PNG (Instagram 1:1)
"""

import math, numpy as np
import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
import matplotlib.patches as mpatches
from matplotlib.patches import FancyBboxPatch
from matplotlib.gridspec import GridSpec
from scipy import stats
from datetime import datetime, timedelta
from PIL import Image
from matplotlib.backends.backend_agg import FigureCanvasAgg

# ── Palette ───────────────────────────────────────────────────────────────────
BG     = '#000000'
CARD   = '#0d0d0d'
BORDER = '#222222'
LIU_C  = '#00B0F0'
DJ_C   = '#ffce00'
WHITE  = '#ffffff'
LGRAY  = '#A0A0A0'
DGRAY  = '#363636'

# ── Sinclair (exact formula from backend/sinclair/sinclair.go) ────────────────
_COEFFS = {
    2001: (0.938573813, 135.390),
    2005: (0.845716976, 168.091),
    2009: (0.784780654, 173.961),
    2013: (0.794358141, 174.393),
    2017: (0.75194503,  175.508),
    2021: (0.722762521, 193.609),
}
def _coeff(year):
    if   year <= 2004: return _COEFFS[2001]
    elif year <= 2008: return _COEFFS[2005]
    elif year <= 2012: return _COEFFS[2009]
    elif year <= 2016: return _COEFFS[2013]
    elif year <= 2020: return _COEFFS[2017]
    else:              return _COEFFS[2021]

def sinclair(bw, total, year):
    if total <= 0 or bw <= 20: return 0.0
    A, b = _coeff(year)
    if bw < b:
        X = math.log10(bw / b)
        return total * 10**(A * X**2)
    return float(total)

# ── Raw data ──────────────────────────────────────────────────────────────────
LIU_RAW = [
    ('2022-12-05', '2022 IWF WC',    '89kg',   88.50, [160, 166,-171], [205, 211, 215], 381),
    ('2023-05-05', '2023 Asian Ch.', '96kg',   89.43, [-170,170, 175], [210,-223,-223], 385),
    ('2023-09-04', '2023 IWF WC',    '102kg',  98.52, [171, 176, 180], [-215,221, 224], 404),
    ('2023-09-30', 'Asian Games',    '109kg', 100.80, [-175,180, 185], [215, 227, 233], 418),
    ('2023-12-04', '2023 IWF GP II', '102kg', 100.18, [170,-176, 176], [210, 222,-225], 398),
    ('2024-03-31', 'WC Paris Qual.', '102kg', 101.84, [175, 181,-186], [220, 225, 232], 413),
    ('2024-08-07', 'Paris Olympics', '102kg', 101.75, [178, 183, 186], [220,-228,-233], 406),
    ('2025-05-09', '2025 Asian Ch.', '102kg', 101.49, [171, 180,-183], [220, 230,-234], 410),
]
DJ_RAW = [
    ('2017-06-15', '2017 Jr WC',      '105kg', 100.30, [155, 159, 162], [-185,185, 190], 352),
    ('2017-11-27', '2017 IWF WC',     '105kg', 103.24, [164, 169, 174], [194, 199,-203], 373),
    ('2018-04-20', '2018 Asian Jr.',  '105kg', 102.20, [160, 166, 170], [191, 197, 202], 372),
    ('2018-07-07', '2018 Jr WC',      '105kg', 102.35, [167,-172,-172], [195, 202,-210], 369),
    ('2018-11-01', '2018 IWF WC',     '102kg', 101.67, [173, 178, 180], [200, 207, 212], 392),
    ('2018-12-19', '5th Qatar Cup',   '109kg', 104.20, [173,-178, 179], [200, 207, 213], 392),
    ('2019-04-18', '2019 Asian Ch.',  '109kg', 108.01, [176, 181, 185], [215, 219, 225], 410),
    ('2019-06-01', '2019 Jr WC',      '109kg', 108.55, [173, 177, 182], [212, 216,   0], 398),
    ('2019-09-18', '2019 IWF WC',     '109kg', 108.60, [-183,184, 188], [221, 226, 229], 417),
    ('2019-12-10', 'IWF World Cup',   '109kg', 108.10, [176, 181,-184], [215, 219,-229], 400),
    ('2019-12-19', '6th Qatar Cup',   '109kg', 108.70, [175, 181, 185], [215, 220,   0], 405),
    ('2020-02-08', '6th Solidarity',  '109kg', 108.65, [177, 182, 189], [-221,221,   0], 410),
    ('2021-04-17', '2020 Asian Ch.',  '109kg', 108.50, [188, 194,-197], [225, 234,-238], 428),
    ('2021-07-23', 'Tokyo Olympics',  '109kg', 109.00, [-189,189, 193], [227,-234, 237], 430),
    ('2021-12-07', '2021 IWF WC',     '109kg', 108.90, [187, 192, 195], [226, 232, 238], 433),
    ('2022-08-11', 'Islamic Sol. G.', '+109kg',121.60, [-190,190, 200], [231, 242, 246], 446),
    ('2023-05-05', '2023 Asian Ch.',  '+109kg',126.76, [189,-195, 195], [230,-240, 242], 437),
    ('2023-09-04', '2023 IWF WC',     '109kg', 108.84, [182,-189, 189], [220, 226,-231], 415),
    ('2023-09-30', 'Asian Games',     '109kg', 109.00, [180, 184, 189], [-220,222, 228], 417),
    ('2024-02-03', '2024 Asian Ch.',  '102kg', 101.70, [175, 180,-183], [214,-219, 220], 400),
    ('2024-03-31', 'WC Paris Qual.',  '109kg', 108.50, [180, 185, 189], [220, 227,   0], 416),
    ('2024-08-07', 'Paris Olympics',  '102kg', 102.00, [180, 185,-189], [219,-224,-232], 404),
    ('2025-05-09', '2025 Asian Ch.',  '109kg', 108.91, [180, 183,-189], [217, 223,   0], 406),
    ('2025-10-02', '2025 IWF WC',     '110kg', 109.85, [189, 193, 196], [227, 232,-245], 428),
]

# ── Derived series ────────────────────────────────────────────────────────────
def parse(raw):
    dates  = [datetime.strptime(r[0],'%Y-%m-%d') for r in raw]
    totals = np.array([r[6] for r in raw], dtype=float)
    sincs  = np.array([sinclair(r[3], r[6], int(r[0][:4])) for r in raw])
    return dates, totals, sincs

def linreg(dates, vals):
    x = np.array([(d - dates[0]).days for d in dates], dtype=float)
    s, b, r, _, _ = stats.linregress(x, vals)
    return s, b, r**2, x

def attempt_rates(raw):
    made = [0]*6; total = [0]*6
    for r in raw:
        for i, v in enumerate(r[4]):
            if v != 0: total[i] += 1; made[i] += int(v > 0)
        for i, v in enumerate(r[5]):
            if v != 0: total[i+3] += 1; made[i+3] += int(v > 0)
    return [m/t if t else 0 for m, t in zip(made, total)]

liu_dates, liu_tot, liu_sin = parse(LIU_RAW)
dj_dates,  dj_tot,  dj_sin  = parse(DJ_RAW)

# regression on Sinclair
liu_ss, liu_sb, liu_sr2, liu_sx = linreg(liu_dates, liu_sin)
dj_ss,  dj_sb,  dj_sr2,  dj_sx  = linreg(dj_dates,  dj_sin)

# 102 kg class
liu_102_sin = np.array([sinclair(r[3],r[6],int(r[0][:4])) for r in LIU_RAW if r[2]=='102kg'])
dj_102_sin  = np.array([sinclair(r[3],r[6],int(r[0][:4])) for r in DJ_RAW  if r[2]=='102kg'])
liu_102_tot = np.array([r[6] for r in LIU_RAW if r[2]=='102kg'])
dj_102_tot  = np.array([r[6] for r in DJ_RAW  if r[2]=='102kg'])

t_sin, p_sin = stats.ttest_ind(liu_102_sin, dj_102_sin, equal_var=False)
pool_sd = np.sqrt((np.std(liu_102_sin,ddof=1)**2 + np.std(dj_102_sin,ddof=1)**2)/2)
d_sin   = (np.mean(liu_102_sin) - np.mean(dj_102_sin)) / pool_sd

liu_rates = attempt_rates(LIU_RAW)
dj_rates  = attempt_rates(DJ_RAW)

print(f"Liu Sinclair:  mean={np.mean(liu_sin):.1f}  SD={np.std(liu_sin,ddof=1):.1f}  best={max(liu_sin):.1f}")
print(f"DJ  Sinclair:  mean={np.mean(dj_sin):.1f}  SD={np.std(dj_sin,ddof=1):.1f}  best={max(dj_sin):.1f}")
print(f"Liu 102kg Sin: mean={np.mean(liu_102_sin):.1f}  SD={np.std(liu_102_sin,ddof=1):.1f}")
print(f"DJ  102kg Sin: mean={np.mean(dj_102_sin):.1f}   SD={np.std(dj_102_sin,ddof=1):.1f}")
print(f"102kg Sinclair: t={t_sin:.2f}  p={p_sin:.3f}  Cohen's d={d_sin:.2f}")
print(f"Liu trend Sin:  {liu_ss*365:+.1f}/yr  R2={liu_sr2:.3f}")
print(f"DJ  trend Sin:  {dj_ss*365:+.1f}/yr  R2={dj_sr2:.3f}")
print(f"Liu 418@100.8  = Sinclair {sinclair(100.80,418,2023):.1f}")
print(f"DJ  446@121.6  = Sinclair {sinclair(121.60,446,2022):.1f}")
print(f"DJ  433@108.9  = Sinclair {sinclair(108.90,433,2021):.1f}  (DJ best)")

# ── Figure ────────────────────────────────────────────────────────────────────
DPI  = 150
INCH = 1080 / DPI

fig = plt.figure(figsize=(INCH, INCH), dpi=DPI, facecolor=BG)
gs  = GridSpec(4, 4, figure=fig,
               left=0.06, right=0.97, top=0.975, bottom=0.02,
               hspace=0.42, wspace=0.32,
               height_ratios=[0.09, 0.36, 0.26, 0.29])

# ─── HEADER ──────────────────────────────────────────────────────────────────
ax_h = fig.add_subplot(gs[0, :])
ax_h.set_facecolor(BG); ax_h.axis('off')
ax_h.text(0.23, 0.82, 'LIU Huanhua', ha='center', va='center',
          color=LIU_C, fontsize=17, fontweight='black', transform=ax_h.transAxes)
ax_h.text(0.23, 0.22, 'CHN', ha='center', va='center',
          color=LIU_C, fontsize=7, alpha=0.65, transform=ax_h.transAxes)
ax_h.text(0.50, 0.72, 'vs', ha='center', va='center',
          color=DGRAY, fontsize=13, fontweight='bold', transform=ax_h.transAxes)
ax_h.text(0.77, 0.82, 'DJURAEV Akbar', ha='center', va='center',
          color=DJ_C, fontsize=17, fontweight='black', transform=ax_h.transAxes)
ax_h.text(0.77, 0.22, 'UZB', ha='center', va='center',
          color=DJ_C, fontsize=7, alpha=0.65, transform=ax_h.transAxes)
ax_h.text(0.50, 0.04, 'IWF CAREER STATISTICAL ANALYSIS  (TOTAL + SINCLAIR)  ·  OpenWeightlifting',
          ha='center', va='center', color=DGRAY, fontsize=4.5,
          style='italic', transform=ax_h.transAxes)

# ─── ROW 1: Two charts side by side ───────────────────────────────────────────
def scatter_chart(ax, title, liu_y, dj_y, ylabel, annot_fn=None):
    ax.set_facecolor(CARD)
    for sp in ax.spines.values(): sp.set_edgecolor(BORDER)
    ax.scatter(liu_dates, liu_y, color=LIU_C, s=22, zorder=5, alpha=0.90)
    ax.scatter(dj_dates,  dj_y,  color=DJ_C,  s=22, zorder=5, alpha=0.90)
    # regression
    for raw, vals, color in [(LIU_RAW, liu_y, LIU_C), (DJ_RAW, dj_y, DJ_C)]:
        dates_ = [datetime.strptime(r[0],'%Y-%m-%d') for r in raw]
        x = np.array([(d - dates_[0]).days for d in dates_], dtype=float)
        s, b, _, _, _ = stats.linregress(x, vals)
        x_ext = np.array([0., x[-1] + 60])
        t_ext = [dates_[0] + timedelta(days=float(v)) for v in x_ext]
        ax.plot(t_ext, b + s * x_ext, color=color, lw=1.2, ls='--', alpha=0.45, zorder=3)
    # mean lines
    ax.axhline(np.mean(liu_y), color=LIU_C, lw=0.6, ls=':', alpha=0.28)
    ax.axhline(np.mean(dj_y),  color=DJ_C,  lw=0.6, ls=':', alpha=0.28)
    if annot_fn:
        annot_fn(ax)
    ax.set_title(title, color=LGRAY, fontsize=6, pad=3)
    ax.set_ylabel(ylabel, color=LGRAY, fontsize=5.5)
    ax.tick_params(colors=LGRAY, labelsize=5, length=2)
    ax.xaxis.set_tick_params(rotation=30)

def annot_total(ax):
    # H2H lines
    for hd in [datetime(2023,9,30), datetime(2024,8,7)]:
        ax.axvline(hd, color=WHITE, lw=0.6, ls=':', alpha=0.20)
    # Asian Games
    ax.annotate('', xy=(datetime(2023,9,30), 419), xytext=(datetime(2023,9,30), 426),
                arrowprops=dict(arrowstyle='->', color=LIU_C, lw=0.9))
    ax.annotate('', xy=(datetime(2023,9,30), 415), xytext=(datetime(2023,9,30), 408),
                arrowprops=dict(arrowstyle='->', color=DJ_C, lw=0.9))
    ax.text(datetime(2023,9,30), 428, '418 v 417', ha='center', va='bottom',
            color=WHITE, fontsize=4, alpha=0.85)
    # Olympics
    ax.annotate('', xy=(datetime(2024,8,7), 407), xytext=(datetime(2024,8,7), 414),
                arrowprops=dict(arrowstyle='->', color=LIU_C, lw=0.9))
    ax.annotate('', xy=(datetime(2024,8,7), 402), xytext=(datetime(2024,8,7), 395),
                arrowprops=dict(arrowstyle='->', color=DJ_C, lw=0.9))
    ax.text(datetime(2024,8,7), 416, '406 v 404', ha='center', va='bottom',
            color=WHITE, fontsize=4, alpha=0.85)
    ax.text(datetime(2022,8,11)+timedelta(days=35), 448, '*+109kg', color=DJ_C, fontsize=3.8, alpha=0.60)

def annot_sinclair(ax):
    # Liu best
    ax.annotate('', xy=(datetime(2023,9,30), 478.5), xytext=(datetime(2023,9,30), 485),
                arrowprops=dict(arrowstyle='->', color=LIU_C, lw=0.9))
    ax.text(datetime(2023,9,30), 486, '477.8', ha='center', va='bottom',
            color=LIU_C, fontsize=4.2, fontweight='bold')
    # DJ best (2021 WC)
    ax.annotate('', xy=(datetime(2021,12,7), 481), xytext=(datetime(2021,12,7), 488),
                arrowprops=dict(arrowstyle='->', color=DJ_C, lw=0.9))
    ax.text(datetime(2021,12,7), 489, '480.4', ha='center', va='bottom',
            color=DJ_C, fontsize=4.2, fontweight='bold')
    # equivalence note
    ax.text(datetime(2022,7,1), 426, 'DJ 446 kg @ +109\n= 477.3 Sinclair',
            ha='center', va='center', color=DJ_C, fontsize=3.8, alpha=0.70,
            linespacing=1.4)
    ax.text(0.02, 0.06,
            f'Liu trend:  {liu_ss*365:+.1f}/yr   R²={liu_sr2:.2f}\n'
            f'DJ trend:  {dj_ss*365:+.1f}/yr   R²={dj_sr2:.2f}',
            transform=ax.transAxes, color=LGRAY, fontsize=4.8, va='bottom', linespacing=1.6)

ax_t = fig.add_subplot(gs[1, :2])
ax_s = fig.add_subplot(gs[1, 2:])

scatter_chart(ax_t, 'Raw Total (kg)', liu_tot, dj_tot, 'Total (kg)', annot_total)
scatter_chart(ax_s, 'Sinclair Score  (bodyweight-adjusted)', liu_sin, dj_sin, 'Sinclair', annot_sinclair)

liu_p = mpatches.Patch(color=LIU_C, label='LIU')
dj_p  = mpatches.Patch(color=DJ_C,  label='DJ')
ax_t.legend(handles=[liu_p, dj_p], loc='upper left', facecolor=CARD,
            edgecolor=BORDER, labelcolor=WHITE, fontsize=5.5, framealpha=0.88,
            handlelength=0.9, borderpad=0.5)

# ─── ROW 2: Stats cards ───────────────────────────────────────────────────────
def card(ax, title):
    ax.set_facecolor(CARD)
    for sp in ax.spines.values(): sp.set_edgecolor(BORDER)
    ax.axis('off')
    ax.text(0.5, 0.95, title, ha='center', va='top', color=WHITE,
            fontsize=6.5, fontweight='bold', transform=ax.transAxes)

def srow(ax, y, lv, mid, rv, lc=LGRAY, rc=LGRAY, lb=False, rb=False):
    ax.text(0.05, y, lv, ha='left',   va='center', color=lc, fontsize=5.8,
            fontweight='bold' if lb else 'normal', transform=ax.transAxes)
    ax.text(0.50, y, mid, ha='center', va='center', color=LGRAY, fontsize=4.8,
            transform=ax.transAxes)
    ax.text(0.95, y, rv, ha='right',  va='center', color=rc, fontsize=5.8,
            fontweight='bold' if rb else 'normal', transform=ax.transAxes)

# Card A: Sinclair comparison
a = fig.add_subplot(gs[2, :2]); card(a, 'SINCLAIR  (bodyweight-adjusted)')
srow(a, 0.78, f'{max(liu_sin):.1f}', 'Career Best', f'{max(dj_sin):.1f}',
     lc=LGRAY, rc=DJ_C, rb=True)
srow(a, 0.60, f'{np.mean(liu_sin):.1f}', 'Career Mean', f'{np.mean(dj_sin):.1f}',
     lc=LIU_C, rc=LGRAY, lb=True)
srow(a, 0.42,
     f'{np.mean(liu_102_sin):.1f}', '102 kg Mean Sin.',  f'{np.mean(dj_102_sin):.1f}',
     lc=LIU_C, rc=LGRAY, lb=True)
srow(a, 0.24,
     f'+/-{np.std(liu_102_sin,ddof=1):.1f}', f"102kg  Cohen's d={d_sin:.2f}  p={p_sin:.2f}",
     f'+/-{np.std(dj_102_sin,ddof=1):.1f}',
     lc=LIU_C, rc=LGRAY, lb=True)
a.text(0.5, 0.05,
       'Liu 418 kg @ 100.8 kg BW  =  477.8 Sinclair',
       ha='center', va='bottom', color=LIU_C, fontsize=4.5,
       fontweight='bold', transform=a.transAxes)
a.text(0.5, 0.005,
       'DJ  446 kg @ 121.6 kg BW  =  477.3 Sinclair  (same score, 28 kg more on the bar)',
       ha='center', va='bottom', color=DJ_C, fontsize=4.0,
       transform=a.transAxes)

# Card B: head-to-head
b = fig.add_subplot(gs[2, 2:]); card(b, 'HEAD-TO-HEAD  (same event, same class)')
b.text(0.50, 0.72, '2  -  0', ha='center', va='center',
       color=LIU_C, fontsize=26, fontweight='black', transform=b.transAxes)
b.text(0.50, 0.54, 'Liu leads all direct meetings',
       ha='center', va='center', color=LGRAY, fontsize=5.5, transform=b.transAxes)
for y_ev, y_sc, evname, lt, dt, ls, ds in [
    (0.41, 0.27, 'Asian Games 2023  (109 kg)',    418, 417, 477.8, 462.5),
    (0.14, 0.00, 'Paris Olympics 2024  (102 kg)', 406, 404, 462.3, 459.6),
]:
    b.text(0.50, y_ev, evname, ha='center', va='center',
           color=LGRAY, fontsize=4.8, transform=b.transAxes)
    b.text(0.10, y_sc, str(lt), ha='center', va='center',
           color=LIU_C, fontsize=7, fontweight='bold', transform=b.transAxes)
    b.text(0.50, y_sc+0.10, f'+{lt-dt} kg', ha='center', va='center',
           color=DGRAY, fontsize=4.5, transform=b.transAxes)
    b.text(0.90, y_sc, str(dt), ha='center', va='center',
           color=LGRAY, fontsize=7, transform=b.transAxes)
    b.text(0.10, y_sc-0.10, f'S {ls:.0f}', ha='center', va='center',
           color=LIU_C, fontsize=4, alpha=0.7, transform=b.transAxes)
    b.text(0.90, y_sc-0.10, f'S {ds:.0f}', ha='center', va='center',
           color=LGRAY, fontsize=4, alpha=0.7, transform=b.transAxes)

# ─── ROW 3: Attempt success rates ────────────────────────────────────────────
ax_r = fig.add_subplot(gs[3, :])
ax_r.set_facecolor(CARD)
for sp in ax_r.spines.values(): sp.set_edgecolor(BORDER)
ax_r.axis('off')

ax_r.text(0.50, 0.97, 'ATTEMPT SUCCESS RATES',
          ha='center', va='top', color=WHITE,
          fontsize=7, fontweight='bold', transform=ax_r.transAxes)

labels  = ['S1', 'S2', 'S3', 'CJ1', 'CJ2', 'CJ3']
X_LEFT  = 0.01; X_RIGHT = 0.53
BAR_MAX = 0.38
BAR_H   = 0.088; ROW_H = 0.22; ROW_GAP = 0.05
row_tops = [0.86, 0.86-(ROW_H+ROW_GAP), 0.86-2*(ROW_H+ROW_GAP)]

for col in range(2):
    x0 = X_LEFT if col == 0 else X_RIGHT
    sec = 'SNATCH' if col == 0 else 'C&J'
    ax_r.text(x0 + BAR_MAX*0.5, 0.95, sec, ha='center', va='top',
              color=DGRAY, fontsize=5.5, fontweight='bold', transform=ax_r.transAxes)
    for row_i in range(3):
        idx = col*3 + row_i
        lr  = liu_rates[idx]; dr = dj_rates[idx]
        rt  = row_tops[row_i]
        y_dj  = rt - BAR_H*0.05
        y_liu = rt - BAR_H - BAR_H*0.15
        for y_bar, rate, color in [(y_dj, dr, DJ_C), (y_liu, lr, LIU_C)]:
            ax_r.add_patch(FancyBboxPatch((x0, y_bar-BAR_H), BAR_MAX, BAR_H,
                boxstyle='round,pad=0.002', facecolor=DGRAY, alpha=0.4,
                transform=ax_r.transAxes, zorder=1))
            if rate > 0:
                ax_r.add_patch(FancyBboxPatch((x0, y_bar-BAR_H), BAR_MAX*rate, BAR_H,
                    boxstyle='round,pad=0.002', facecolor=color, alpha=0.88,
                    transform=ax_r.transAxes, zorder=2))
        ax_r.text(x0-0.005, rt-BAR_H-BAR_H*0.4, labels[idx], ha='right', va='center',
                  color=WHITE, fontsize=6, fontweight='bold', transform=ax_r.transAxes)
        px = x0 + BAR_MAX + 0.008
        ax_r.text(px, y_dj  - BAR_H*0.5, f'{dr*100:.0f}%', ha='left', va='center',
                  color=DJ_C,  fontsize=5.2, transform=ax_r.transAxes)
        ax_r.text(px, y_liu - BAR_H*0.5, f'{lr*100:.0f}%', ha='left', va='center',
                  color=LIU_C, fontsize=5.2, transform=ax_r.transAxes)

# legend
for x_leg, color, name in [(0.74, DJ_C, 'DJURAEV'), (0.88, LIU_C, 'LIU')]:
    ax_r.add_patch(FancyBboxPatch((x_leg, 0.90), 0.018, 0.07,
        boxstyle='round,pad=0.002', facecolor=color, alpha=0.88,
        transform=ax_r.transAxes))
    ax_r.text(x_leg+0.022, 0.935, name, ha='left', va='center',
              color=color, fontsize=5.2, transform=ax_r.transAxes)

# ── Save: exact 1080×1080 ─────────────────────────────────────────────────────
out = '/home/user/OpenWeightlifting/liu_akbar_comparison.png'
canvas = FigureCanvasAgg(fig)
canvas.draw()
buf = canvas.buffer_rgba()
img = Image.frombuffer('RGBA', canvas.get_width_height(), buf, 'raw', 'RGBA', 0, 1)
img = img.convert('RGB')
cw, ch = img.size; target = 1080
if cw != target or ch != target:
    final = Image.new('RGB', (target, target), (0,0,0))
    final.paste(img, ((target-cw)//2, (target-ch)//2))
    img = final
img.save(out, dpi=(target, target))
plt.close()
print(f'\nSaved {img.size[0]}x{img.size[1]}px -> {out}')
