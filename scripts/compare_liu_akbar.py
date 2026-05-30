"""
Statistical comparison infographic: LIU Huanhua vs DJURAEV Akbar
Output: exactly 1080×1080 PNG (Instagram 1:1)
"""

import numpy as np
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

# ── Raw data ──────────────────────────────────────────────────────────────────
LIU_RAW = [
    ('2022-12-05', '2022 IWF WC',     '89kg',   88.50, [160, 166,-171], [205, 211, 215], 381),
    ('2023-05-05', '2023 Asian Ch.',  '96kg',   89.43, [-170,170, 175], [210,-223,-223], 385),
    ('2023-09-04', '2023 IWF WC',     '102kg',  98.52, [171, 176, 180], [-215,221, 224], 404),
    ('2023-09-30', 'Asian Games',     '109kg', 100.80, [-175,180, 185], [215, 227, 233], 418),
    ('2023-12-04', '2023 IWF GP II',  '102kg', 100.18, [170,-176, 176], [210, 222,-225], 398),
    ('2024-03-31', 'WC Paris Qual.',  '102kg', 101.84, [175, 181,-186], [220, 225, 232], 413),
    ('2024-08-07', 'Paris Olympics',  '102kg', 101.75, [178, 183, 186], [220,-228,-233], 406),
    ('2025-05-09', '2025 Asian Ch.',  '102kg', 101.49, [171, 180,-183], [220, 230,-234], 410),
]
DJ_RAW = [
    ('2017-06-15', '2017 Jr WC',       '105kg', 100.30, [155, 159, 162], [-185,185, 190], 352),
    ('2017-11-27', '2017 IWF WC',      '105kg', 103.24, [164, 169, 174], [194, 199,-203], 373),
    ('2018-04-20', '2018 Asian Jr.',   '105kg', 102.20, [160, 166, 170], [191, 197, 202], 372),
    ('2018-07-07', '2018 Jr WC',       '105kg', 102.35, [167,-172,-172], [195, 202,-210], 369),
    ('2018-11-01', '2018 IWF WC',      '102kg', 101.67, [173, 178, 180], [200, 207, 212], 392),
    ('2018-12-19', '5th Qatar Cup',    '109kg', 104.20, [173,-178, 179], [200, 207, 213], 392),
    ('2019-04-18', '2019 Asian Ch.',   '109kg', 108.01, [176, 181, 185], [215, 219, 225], 410),
    ('2019-06-01', '2019 Jr WC',       '109kg', 108.55, [173, 177, 182], [212, 216,   0], 398),
    ('2019-09-18', '2019 IWF WC',      '109kg', 108.60, [-183,184, 188], [221, 226, 229], 417),
    ('2019-12-10', 'IWF World Cup',    '109kg', 108.10, [176, 181,-184], [215, 219,-229], 400),
    ('2019-12-19', '6th Qatar Cup',    '109kg', 108.70, [175, 181, 185], [215, 220,   0], 405),
    ('2020-02-08', '6th Solidarity',   '109kg', 108.65, [177, 182, 189], [-221,221,   0], 410),
    ('2021-04-17', '2020 Asian Ch.',   '109kg', 108.50, [188, 194,-197], [225, 234,-238], 428),
    ('2021-07-23', 'Tokyo Olympics',   '109kg', 109.00, [-189,189, 193], [227,-234, 237], 430),
    ('2021-12-07', '2021 IWF WC',      '109kg', 108.90, [187, 192, 195], [226, 232, 238], 433),
    ('2022-08-11', 'Islamic Sol. G.',  '+109kg',121.60, [-190,190, 200], [231, 242, 246], 446),
    ('2023-05-05', '2023 Asian Ch.',   '+109kg',126.76, [189,-195, 195], [230,-240, 242], 437),
    ('2023-09-04', '2023 IWF WC',      '109kg', 108.84, [182,-189, 189], [220, 226,-231], 415),
    ('2023-09-30', 'Asian Games',      '109kg', 109.00, [180, 184, 189], [-220,222, 228], 417),
    ('2024-02-03', '2024 Asian Ch.',   '102kg', 101.70, [175, 180,-183], [214,-219, 220], 400),
    ('2024-03-31', 'WC Paris Qual.',   '109kg', 108.50, [180, 185, 189], [220, 227,   0], 416),
    ('2024-08-07', 'Paris Olympics',   '102kg', 102.00, [180, 185,-189], [219,-224,-232], 404),
    ('2025-05-09', '2025 Asian Ch.',   '109kg', 108.91, [180, 183,-189], [217, 223,   0], 406),
    ('2025-10-02', '2025 IWF WC',      '110kg', 109.85, [189, 193, 196], [227, 232,-245], 428),
]

# ── Stats helpers ─────────────────────────────────────────────────────────────
def attempt_rates(raw):
    made = [0]*6; total = [0]*6
    for r in raw:
        for i, v in enumerate(r[4]):
            if v != 0: total[i] += 1; made[i] += int(v > 0)
        for i, v in enumerate(r[5]):
            if v != 0: total[i+3] += 1; made[i+3] += int(v > 0)
    return [m/t if t else 0 for m, t in zip(made, total)]

def linreg(raw):
    dates  = [datetime.strptime(r[0], '%Y-%m-%d') for r in raw]
    totals = np.array([r[6] for r in raw], dtype=float)
    x = np.array([(d - dates[0]).days for d in dates], dtype=float)
    s, b, r, _, _ = stats.linregress(x, totals)
    return s, b, r**2, x, dates

# ── Compute ───────────────────────────────────────────────────────────────────
liu_totals = np.array([r[6] for r in LIU_RAW])
dj_totals  = np.array([r[6] for r in DJ_RAW])
liu_dates  = [datetime.strptime(r[0],'%Y-%m-%d') for r in LIU_RAW]
dj_dates   = [datetime.strptime(r[0],'%Y-%m-%d') for r in DJ_RAW]

liu_102 = np.array([r[6] for r in LIU_RAW if r[2]=='102kg'])  # n=5
dj_102  = np.array([r[6] for r in DJ_RAW  if r[2]=='102kg'])  # n=3

t_102, p_102 = stats.ttest_ind(liu_102, dj_102, equal_var=False)
pool_sd = np.sqrt((np.std(liu_102,ddof=1)**2 + np.std(dj_102,ddof=1)**2)/2)
d_102   = (np.mean(liu_102) - np.mean(dj_102)) / pool_sd

liu_rates = attempt_rates(LIU_RAW)
dj_rates  = attempt_rates(DJ_RAW)

liu_s, liu_b, liu_r2, liu_x, _ = linreg(LIU_RAW)
dj_s,  dj_b,  dj_r2,  dj_x, _ = linreg(DJ_RAW)

print(f"Liu 102kg: n={len(liu_102)}, mean={np.mean(liu_102):.1f}, SD={np.std(liu_102,ddof=1):.1f}")
print(f"DJ  102kg: n={len(dj_102)},  mean={np.mean(dj_102):.1f},  SD={np.std(dj_102,ddof=1):.1f}")
print(f"Welch t={t_102:.2f}  p={p_102:.3f}  Cohen's d={d_102:.2f}")
print(f"Liu trend: {liu_s*365:+.1f} kg/yr  R2={liu_r2:.3f}")
print(f"DJ  trend: {dj_s*365:+.1f} kg/yr  R2={dj_r2:.3f}")
print(f"Liu rates: S={[f'{r*100:.0f}%' for r in liu_rates[:3]]}  CJ={[f'{r*100:.0f}%' for r in liu_rates[3:]]}")
print(f"DJ  rates: S={[f'{r*100:.0f}%' for r in dj_rates[:3]]}  CJ={[f'{r*100:.0f}%' for r in dj_rates[3:]]}")

# ── Figure ────────────────────────────────────────────────────────────────────
DPI = 150
INCH = 1080 / DPI

fig = plt.figure(figsize=(INCH, INCH), dpi=DPI, facecolor=BG)
gs  = GridSpec(4, 2, figure=fig,
               left=0.06, right=0.97, top=0.975, bottom=0.02,
               hspace=0.35, wspace=0.26,
               height_ratios=[0.09, 0.40, 0.24, 0.27])

# ─── HEADER ──────────────────────────────────────────────────────────────────
ax_h = fig.add_subplot(gs[0, :])
ax_h.set_facecolor(BG); ax_h.axis('off')
ax_h.text(0.23, 0.82, 'LIU Huanhua', ha='center', va='center',
          color=LIU_C, fontsize=17, fontweight='black', transform=ax_h.transAxes)
ax_h.text(0.23, 0.20, 'CHN', ha='center', va='center',
          color=LIU_C, fontsize=7, alpha=0.65, transform=ax_h.transAxes)
ax_h.text(0.50, 0.72, 'vs', ha='center', va='center',
          color=DGRAY, fontsize=13, fontweight='bold', transform=ax_h.transAxes)
ax_h.text(0.77, 0.82, 'DJURAEV Akbar', ha='center', va='center',
          color=DJ_C, fontsize=17, fontweight='black', transform=ax_h.transAxes)
ax_h.text(0.77, 0.20, 'UZB', ha='center', va='center',
          color=DJ_C, fontsize=7, alpha=0.65, transform=ax_h.transAxes)
ax_h.text(0.50, 0.04, 'IWF CAREER STATISTICAL ANALYSIS  ·  OpenWeightlifting',
          ha='center', va='center', color=DGRAY, fontsize=4.5,
          style='italic', transform=ax_h.transAxes)

# ─── MAIN CHART: Totals scatter + regression ──────────────────────────────────
ax_m = fig.add_subplot(gs[1, :])
ax_m.set_facecolor(CARD)
for sp in ax_m.spines.values(): sp.set_edgecolor(BORDER)

ax_m.scatter(liu_dates, liu_totals, color=LIU_C, s=30, zorder=5, alpha=0.90)
ax_m.scatter(dj_dates,  dj_totals,  color=DJ_C,  s=30, zorder=5, alpha=0.90)

# regression lines
for raw, slope, intercept, x_arr, color in [
    (LIU_RAW, liu_s, liu_b, liu_x, LIU_C),
    (DJ_RAW,  dj_s,  dj_b,  dj_x,  DJ_C),
]:
    base = datetime.strptime(raw[0][0], '%Y-%m-%d')
    x_ext = np.array([0., x_arr[-1] + 60])
    t_ext = [base + timedelta(days=float(v)) for v in x_ext]
    ax_m.plot(t_ext, intercept + slope * x_ext,
              color=color, lw=1.6, ls='--', alpha=0.50, zorder=3)

# mean reference lines
ax_m.axhline(np.mean(liu_totals), color=LIU_C, lw=0.7, ls=':', alpha=0.28, zorder=2)
ax_m.axhline(np.mean(dj_totals),  color=DJ_C,  lw=0.7, ls=':', alpha=0.28, zorder=2)

# H2H vertical guides
for hd in [datetime(2023,9,30), datetime(2024,8,7)]:
    ax_m.axvline(hd, color=WHITE, lw=0.7, ls=':', alpha=0.22, zorder=1)

# Asian Games annotation
ax_m.annotate('', xy=(datetime(2023,9,30), 419.5),
              xytext=(datetime(2023,9,30), 427),
              arrowprops=dict(arrowstyle='->', color=LIU_C, lw=1.1, alpha=0.9))
ax_m.annotate('', xy=(datetime(2023,9,30), 414.5),
              xytext=(datetime(2023,9,30), 407),
              arrowprops=dict(arrowstyle='->', color=DJ_C, lw=1.1, alpha=0.9))
ax_m.text(datetime(2023,9,30), 429.5, 'Asian Games\n418 vs 417',
          ha='center', va='bottom', color=WHITE, fontsize=4.8, alpha=0.9,
          linespacing=1.5)

# Olympics annotation
ax_m.annotate('', xy=(datetime(2024,8,7), 407.5),
              xytext=(datetime(2024,8,7), 415),
              arrowprops=dict(arrowstyle='->', color=LIU_C, lw=1.1, alpha=0.9))
ax_m.annotate('', xy=(datetime(2024,8,7), 401.5),
              xytext=(datetime(2024,8,7), 394),
              arrowprops=dict(arrowstyle='->', color=DJ_C, lw=1.1, alpha=0.9))
ax_m.text(datetime(2024,8,7), 417, 'Paris Olympics\n406 vs 404',
          ha='center', va='bottom', color=WHITE, fontsize=4.8, alpha=0.9,
          linespacing=1.5)

# +109 kg asterisk
ax_m.text(datetime(2022,8,11) + timedelta(days=50), 447.5,
          '* at +109 kg (BW 121 kg)', color=DJ_C, fontsize=4.3, alpha=0.60)

ax_m.text(0.01, 0.04,
          f'Trend  Liu: {liu_s*365:+.1f} kg/yr  R²={liu_r2:.2f}    '
          f'Djuraev: {dj_s*365:+.1f} kg/yr  R²={dj_r2:.2f}',
          transform=ax_m.transAxes, color=LGRAY, fontsize=5.2, va='bottom')

ax_m.set_ylabel('Total (kg)', color=LGRAY, fontsize=6)
ax_m.tick_params(colors=LGRAY, labelsize=5.5, length=2.5)
ax_m.xaxis.set_tick_params(rotation=30)

liu_patch = mpatches.Patch(color=LIU_C, label='LIU Huanhua')
dj_patch  = mpatches.Patch(color=DJ_C,  label='DJURAEV Akbar')
ax_m.legend(handles=[liu_patch, dj_patch], loc='upper left',
            facecolor=CARD, edgecolor=BORDER, labelcolor=WHITE,
            fontsize=5.8, framealpha=0.88, handlelength=1)

# ─── ROW 2: Stats cards ───────────────────────────────────────────────────────
def card(ax, title):
    ax.set_facecolor(CARD)
    for sp in ax.spines.values(): sp.set_edgecolor(BORDER)
    ax.axis('off')
    ax.text(0.5, 0.95, title, ha='center', va='top', color=WHITE,
            fontsize=7, fontweight='bold', transform=ax.transAxes)

def stat_row(ax, y, lval, label, rval, lcolor=LGRAY, rcolor=LGRAY, lbold=False, rbold=False):
    ax.text(0.05, y, lval,  ha='left',   va='center', color=lcolor,
            fontsize=6, fontweight='bold' if lbold else 'normal',
            transform=ax.transAxes)
    ax.text(0.50, y, label, ha='center', va='center', color=LGRAY,
            fontsize=5, transform=ax.transAxes)
    ax.text(0.95, y, rval,  ha='right',  va='center', color=rcolor,
            fontsize=6, fontweight='bold' if rbold else 'normal',
            transform=ax.transAxes)

# Card A: 102 kg class comparison
a = fig.add_subplot(gs[2, 0]); card(a, '@ 102 KG CLASS  (direct comparison)')
liu_m = np.mean(liu_102); liu_sd = np.std(liu_102, ddof=1); liu_cv = liu_sd/liu_m*100
dj_m  = np.mean(dj_102);  dj_sd  = np.std(dj_102,  ddof=1); dj_cv  = dj_sd /dj_m *100
stat_row(a, 0.76, f'{liu_m:.1f} kg', 'Mean Total',  f'{dj_m:.1f} kg',
         lcolor=LIU_C, rcolor=LGRAY, lbold=True)
stat_row(a, 0.58, f'+/-{liu_sd:.1f}', 'Std Dev (kg)', f'+/-{dj_sd:.1f}',
         lcolor=LIU_C, rcolor=LGRAY, lbold=True)
stat_row(a, 0.40, f'{liu_cv:.1f}%', 'CV %',  f'{dj_cv:.1f}%',
         lcolor=LIU_C, rcolor=LGRAY, lbold=True)
stat_row(a, 0.22, f'n={len(liu_102)}', f"Cohen's d={d_102:.2f}  p={p_102:.2f}",
         f'n={len(dj_102)}', lcolor=LGRAY, rcolor=LGRAY)
a.text(0.50, 0.04, 'Liu: lower variance, higher mean at 102 kg',
       ha='center', va='bottom', color=LIU_C, fontsize=4.2,
       fontweight='bold', transform=a.transAxes)

# Card B: head-to-head
b = fig.add_subplot(gs[2, 1]); card(b, 'HEAD-TO-HEAD  (same event, same class)')
b.text(0.50, 0.72, '2  -  0', ha='center', va='center',
       color=LIU_C, fontsize=26, fontweight='black', transform=b.transAxes)
b.text(0.50, 0.54, 'Liu leads all direct meetings',
       ha='center', va='center', color=LGRAY, fontsize=5.5, transform=b.transAxes)
for y_ev, y_sc, evname, lt, dt in [
    (0.41, 0.27, 'Asian Games 2023  (109 kg)',    418, 417),
    (0.16, 0.03, 'Paris Olympics 2024  (102 kg)', 406, 404),
]:
    b.text(0.50, y_ev, evname, ha='center', va='center',
           color=LGRAY, fontsize=4.8, transform=b.transAxes)
    b.text(0.12, y_sc, str(lt), ha='center', va='center',
           color=LIU_C, fontsize=8, fontweight='bold', transform=b.transAxes)
    b.text(0.50, y_sc, f'+{lt-dt} kg', ha='center', va='center',
           color=DGRAY, fontsize=5.2, transform=b.transAxes)
    b.text(0.88, y_sc, str(dt), ha='center', va='center',
           color=LGRAY, fontsize=8, transform=b.transAxes)

# ─── ROW 3: Attempt success rate bars ────────────────────────────────────────
ax_r = fig.add_subplot(gs[3, :])
ax_r.set_facecolor(CARD)
for sp in ax_r.spines.values(): sp.set_edgecolor(BORDER)
ax_r.axis('off')

ax_r.text(0.50, 0.97, 'ATTEMPT SUCCESS RATES',
          ha='center', va='top', color=WHITE,
          fontsize=7, fontweight='bold', transform=ax_r.transAxes)

# 6 attempt slots: 2 columns × 3 rows, side by side
# Axes coords (0-1 × 0-1) for the whole bar section
labels = ['S1', 'S2', 'S3', 'CJ1', 'CJ2', 'CJ3']
COL_W  = 0.46    # width of each column group in axes coords
X_LEFT = 0.01    # x start of snatch column
X_RIGHT= 0.53    # x start of CJ column
BAR_MAX= 0.38    # max bar width in axes coords

ROW_H  = 0.22    # total height of one attempt group (two bars)
BAR_H  = 0.088   # height of each individual bar in axes coords
ROW_GAP= 0.05    # gap between rows

# 3 row centers (from top downward in axes coords, leaving room for header)
row_tops = [0.86, 0.86 - (ROW_H + ROW_GAP), 0.86 - 2*(ROW_H + ROW_GAP)]

for col in range(2):    # 0=snatch, 1=CJ
    x0 = X_LEFT if col == 0 else X_RIGHT
    section = 'SNATCH' if col == 0 else 'C&J'
    ax_r.text(x0 + BAR_MAX * 0.5, 0.95, section,
              ha='center', va='top', color=DGRAY, fontsize=5.5,
              fontweight='bold', transform=ax_r.transAxes)

    for row_i in range(3):
        attempt_idx = col * 3 + row_i
        lbl    = labels[attempt_idx]
        lr     = liu_rates[attempt_idx]
        dr     = dj_rates[attempt_idx]
        rt     = row_tops[row_i]

        y_dj  = rt - BAR_H * 0.05           # DJ bar top position
        y_liu = rt - BAR_H - BAR_H * 0.15   # Liu bar top position (below DJ)

        # BG tracks
        ax_r.add_patch(FancyBboxPatch(
            (x0, y_dj - BAR_H), BAR_MAX, BAR_H,
            boxstyle='round,pad=0.002', facecolor=DGRAY, alpha=0.4,
            transform=ax_r.transAxes, zorder=1))
        ax_r.add_patch(FancyBboxPatch(
            (x0, y_liu - BAR_H), BAR_MAX, BAR_H,
            boxstyle='round,pad=0.002', facecolor=DGRAY, alpha=0.4,
            transform=ax_r.transAxes, zorder=1))

        # Filled bars
        if dr > 0:
            ax_r.add_patch(FancyBboxPatch(
                (x0, y_dj - BAR_H), BAR_MAX * dr, BAR_H,
                boxstyle='round,pad=0.002', facecolor=DJ_C, alpha=0.88,
                transform=ax_r.transAxes, zorder=2))
        if lr > 0:
            ax_r.add_patch(FancyBboxPatch(
                (x0, y_liu - BAR_H), BAR_MAX * lr, BAR_H,
                boxstyle='round,pad=0.002', facecolor=LIU_C, alpha=0.88,
                transform=ax_r.transAxes, zorder=2))

        # Attempt label (left)
        label_x = x0 - 0.005
        ax_r.text(label_x, rt - BAR_H - BAR_H * 0.4,
                  lbl, ha='right', va='center', color=WHITE,
                  fontsize=6, fontweight='bold', transform=ax_r.transAxes)

        # % labels (right)
        pct_x = x0 + BAR_MAX + 0.008
        ax_r.text(pct_x, y_dj  - BAR_H * 0.5, f'{dr*100:.0f}%',
                  ha='left', va='center', color=DJ_C, fontsize=5.2,
                  transform=ax_r.transAxes)
        ax_r.text(pct_x, y_liu - BAR_H * 0.5, f'{lr*100:.0f}%',
                  ha='left', va='center', color=LIU_C, fontsize=5.2,
                  transform=ax_r.transAxes)

# Column divider
ax_r.axvline(0.50, color=BORDER, lw=1.0, ymin=0.08, ymax=0.90)

# Legend for bars — far right of title row
for x_leg, color, name in [(0.74, DJ_C, 'DJURAEV'), (0.88, LIU_C, 'LIU')]:
    ax_r.add_patch(FancyBboxPatch(
        (x_leg, 0.90), 0.018, 0.07,
        boxstyle='round,pad=0.002', facecolor=color, alpha=0.88,
        transform=ax_r.transAxes))
    ax_r.text(x_leg + 0.022, 0.935, name, ha='left', va='center',
              color=color, fontsize=5.2, transform=ax_r.transAxes)

# ── Save: exact 1080×1080 ─────────────────────────────────────────────────────
out = '/home/user/OpenWeightlifting/liu_akbar_comparison.png'

canvas = FigureCanvasAgg(fig)
canvas.draw()
buf = canvas.buffer_rgba()
img = Image.frombuffer('RGBA', canvas.get_width_height(), buf, 'raw', 'RGBA', 0, 1)
img = img.convert('RGB')

# Pad or crop to exact 1080×1080
cw, ch = img.size
target = 1080
if cw != target or ch != target:
    final = Image.new('RGB', (target, target), (0, 0, 0))
    final.paste(img, ((target - cw)//2, (target - ch)//2))
    img = final

img.save(out, dpi=(target, target))
plt.close()
print(f'\nSaved {img.size[0]}x{img.size[1]}px -> {out}')
