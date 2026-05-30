"""
3-image Instagram carousel: LIU Huanhua vs DJURAEV Akbar
Image 1 — Hook: H2H record + Sinclair equivalence
Image 2 — Charts: Total & Sinclair career progressions
Image 3 — Stats: Sinclair deep-dive + attempt rates
All outputs: 1080×1080 PNG
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
CARD2  = '#111111'
BORDER = '#222222'
LIU_C  = '#00B0F0'
DJ_C   = '#ffce00'
WHITE  = '#ffffff'
LGRAY  = '#A0A0A0'
DGRAY  = '#363636'
GREEN  = '#598138'
RED    = '#9d3d3d'

# ── Sinclair (exact backend formula) ─────────────────────────────────────────
_COEFFS = {
    2001: (0.938573813, 135.390), 2005: (0.845716976, 168.091),
    2009: (0.784780654, 173.961), 2013: (0.794358141, 174.393),
    2017: (0.75194503,  175.508), 2021: (0.722762521, 193.609),
}
def _coeff(y):
    if   y <= 2004: return _COEFFS[2001]
    elif y <= 2008: return _COEFFS[2005]
    elif y <= 2012: return _COEFFS[2009]
    elif y <= 2016: return _COEFFS[2013]
    elif y <= 2020: return _COEFFS[2017]
    else:           return _COEFFS[2021]

def sinc(bw, total, year):
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

# ── Derived ───────────────────────────────────────────────────────────────────
def parse(raw):
    dates  = [datetime.strptime(r[0],'%Y-%m-%d') for r in raw]
    totals = np.array([r[6] for r in raw], dtype=float)
    sincs  = np.array([sinc(r[3], r[6], int(r[0][:4])) for r in raw])
    return dates, totals, sincs

def linreg(dates, vals):
    x = np.array([(d - dates[0]).days for d in dates], dtype=float)
    s, b, r, _, _ = stats.linregress(x, vals)
    return s, b, r**2, x

def attempt_rates(raw):
    made = [0]*6; tot = [0]*6
    for r in raw:
        for i, v in enumerate(r[4]):
            if v != 0: tot[i] += 1; made[i] += int(v > 0)
        for i, v in enumerate(r[5]):
            if v != 0: tot[i+3] += 1; made[i+3] += int(v > 0)
    return [m/t if t else 0 for m, t in zip(made, tot)]

liu_dates, liu_tot, liu_sin = parse(LIU_RAW)
dj_dates,  dj_tot,  dj_sin  = parse(DJ_RAW)

liu_ts, liu_tb, liu_tr2, liu_tx = linreg(liu_dates, liu_tot)
dj_ts,  dj_tb,  dj_tr2,  dj_tx  = linreg(dj_dates,  dj_tot)
liu_ss, liu_sb, liu_sr2, liu_sx = linreg(liu_dates, liu_sin)
dj_ss,  dj_sb,  dj_sr2,  dj_sx  = linreg(dj_dates,  dj_sin)

liu_102_sin = np.array([sinc(r[3],r[6],int(r[0][:4])) for r in LIU_RAW if r[2]=='102kg'])
dj_102_sin  = np.array([sinc(r[3],r[6],int(r[0][:4])) for r in DJ_RAW  if r[2]=='102kg'])
liu_102_tot = np.array([r[6] for r in LIU_RAW if r[2]=='102kg'])
dj_102_tot  = np.array([r[6] for r in DJ_RAW  if r[2]=='102kg'])

t_102t, p_102t = stats.ttest_ind(liu_102_tot, dj_102_tot, equal_var=False)
t_102s, p_102s = stats.ttest_ind(liu_102_sin, dj_102_sin, equal_var=False)
pool_s = np.sqrt((np.std(liu_102_sin,ddof=1)**2 + np.std(dj_102_sin,ddof=1)**2)/2)
d_102s = (np.mean(liu_102_sin) - np.mean(dj_102_sin)) / pool_s

liu_rates = attempt_rates(LIU_RAW)
dj_rates  = attempt_rates(DJ_RAW)

# ── Shared helpers ────────────────────────────────────────────────────────────
DPI    = 150
TARGET = 1080
INCH   = TARGET / DPI

def save(fig, path):
    canvas = FigureCanvasAgg(fig)
    canvas.draw()
    buf = canvas.buffer_rgba()
    img = Image.frombuffer('RGBA', canvas.get_width_height(), buf, 'raw', 'RGBA', 0, 1)
    img = img.convert('RGB')
    cw, ch = img.size
    if cw != TARGET or ch != TARGET:
        pad = Image.new('RGB', (TARGET, TARGET), (0,0,0))
        pad.paste(img, ((TARGET-cw)//2, (TARGET-ch)//2))
        img = pad
    img.save(path, dpi=(TARGET, TARGET))
    plt.close(fig)
    print(f'Saved {img.size[0]}x{img.size[1]}px  {path}')

def nameplate(ax, n):
    """Page indicator n/3 in the corner."""
    ax.text(0.97, 0.97, f'{n} / 3', ha='right', va='top', color=DGRAY,
            fontsize=6, transform=ax.transAxes)

def header(ax, subtitle=''):
    ax.set_facecolor(BG); ax.axis('off')
    ax.text(0.22, 0.80, 'LIU Huanhua', ha='center', va='center',
            color=LIU_C, fontsize=15, fontweight='black', transform=ax.transAxes)
    ax.text(0.22, 0.18, 'CHN', ha='center', va='center',
            color=LIU_C, fontsize=6.5, alpha=0.65, transform=ax.transAxes)
    ax.text(0.50, 0.70, 'vs', ha='center', va='center',
            color=DGRAY, fontsize=11, fontweight='bold', transform=ax.transAxes)
    ax.text(0.78, 0.80, 'DJURAEV Akbar', ha='center', va='center',
            color=DJ_C, fontsize=15, fontweight='black', transform=ax.transAxes)
    ax.text(0.78, 0.18, 'UZB', ha='center', va='center',
            color=DJ_C, fontsize=6.5, alpha=0.65, transform=ax.transAxes)
    if subtitle:
        ax.text(0.50, 0.04, subtitle, ha='center', va='center',
                color=DGRAY, fontsize=4.5, style='italic', transform=ax.transAxes)

def card_bg(ax):
    ax.set_facecolor(CARD)
    for sp in ax.spines.values(): sp.set_edgecolor(BORDER)
    ax.axis('off')

def card_title(ax, title, y=0.95):
    ax.text(0.5, y, title, ha='center', va='top', color=WHITE,
            fontsize=7, fontweight='bold', transform=ax.transAxes)

def row3(ax, y, lv, mid, rv, lc=LGRAY, rc=LGRAY, lb=False, rb=False):
    ax.text(0.05, y, lv,  ha='left',   va='center', color=lc, fontsize=6.2,
            fontweight='bold' if lb else 'normal', transform=ax.transAxes)
    ax.text(0.50, y, mid, ha='center', va='center', color=LGRAY, fontsize=5,
            transform=ax.transAxes)
    ax.text(0.95, y, rv,  ha='right',  va='center', color=rc, fontsize=6.2,
            fontweight='bold' if rb else 'normal', transform=ax.transAxes)

def scatter_ax(ax, dates_l, vals_l, dates_d, vals_d, ylabel,
               slope_l, int_l, x_l, slope_d, int_d, x_d,
               annot_fn=None):
    ax.set_facecolor(CARD)
    for sp in ax.spines.values(): sp.set_edgecolor(BORDER)
    ax.scatter(dates_l, vals_l, color=LIU_C, s=28, zorder=5, alpha=0.90)
    ax.scatter(dates_d, vals_d, color=DJ_C,  s=28, zorder=5, alpha=0.90)
    base_l = dates_l[0]; base_d = dates_d[0]
    x_e = np.array([0., x_l[-1] + 60])
    ax.plot([base_l + timedelta(days=float(v)) for v in x_e],
            int_l + slope_l * x_e, color=LIU_C, lw=1.4, ls='--', alpha=0.45)
    x_e = np.array([0., x_d[-1] + 60])
    ax.plot([base_d + timedelta(days=float(v)) for v in x_e],
            int_d + slope_d * x_e, color=DJ_C,  lw=1.4, ls='--', alpha=0.45)
    ax.axhline(np.mean(vals_l), color=LIU_C, lw=0.6, ls=':', alpha=0.25)
    ax.axhline(np.mean(vals_d), color=DJ_C,  lw=0.6, ls=':', alpha=0.25)
    if annot_fn: annot_fn(ax)
    ax.set_ylabel(ylabel, color=LGRAY, fontsize=6)
    ax.tick_params(colors=LGRAY, labelsize=5.5, length=2.5)
    ax.xaxis.set_tick_params(rotation=30)

def bars_section(fig, gs_slot, top=0.97):
    """Attempt success rates bars. gs_slot is a GridSpec entry spanning full width."""
    ax = fig.add_subplot(gs_slot)
    card_bg(ax)
    ax.text(0.5, top, 'ATTEMPT SUCCESS RATES',
            ha='center', va='top', color=WHITE,
            fontsize=7.5, fontweight='bold', transform=ax.transAxes)

    labels  = ['S1','S2','S3','CJ1','CJ2','CJ3']
    X_LEFT  = 0.01; X_RIGHT = 0.53; BAR_MAX = 0.38
    BAR_H   = 0.088; ROW_H = 0.20; ROW_GAP = 0.04
    row_tops = [top-0.14, top-0.14-(ROW_H+ROW_GAP), top-0.14-2*(ROW_H+ROW_GAP)]

    for col in range(2):
        x0  = X_LEFT if col == 0 else X_RIGHT
        sec = 'SNATCH' if col == 0 else 'C&J'
        ax.text(x0 + BAR_MAX*0.5, top-0.04, sec,
                ha='center', va='top', color=DGRAY,
                fontsize=5.5, fontweight='bold', transform=ax.transAxes)
        for row_i in range(3):
            idx   = col*3 + row_i
            lr    = liu_rates[idx]; dr = dj_rates[idx]
            rt    = row_tops[row_i]
            y_dj  = rt - BAR_H*0.05
            y_liu = rt - BAR_H - BAR_H*0.15
            for y_bar, rate, color in [(y_dj, dr, DJ_C), (y_liu, lr, LIU_C)]:
                ax.add_patch(FancyBboxPatch(
                    (x0, y_bar-BAR_H), BAR_MAX, BAR_H,
                    boxstyle='round,pad=0.002', facecolor=DGRAY, alpha=0.35,
                    transform=ax.transAxes, zorder=1))
                if rate > 0:
                    ax.add_patch(FancyBboxPatch(
                        (x0, y_bar-BAR_H), BAR_MAX*rate, BAR_H,
                        boxstyle='round,pad=0.002', facecolor=color, alpha=0.88,
                        transform=ax.transAxes, zorder=2))
            ax.text(x0-0.005, rt-BAR_H-BAR_H*0.4, labels[idx],
                    ha='right', va='center', color=WHITE,
                    fontsize=6.5, fontweight='bold', transform=ax.transAxes)
            px = x0 + BAR_MAX + 0.008
            ax.text(px, y_dj -BAR_H*0.5, f'{dr*100:.0f}%', ha='left', va='center',
                    color=DJ_C,  fontsize=5.5, transform=ax.transAxes)
            ax.text(px, y_liu-BAR_H*0.5, f'{lr*100:.0f}%', ha='left', va='center',
                    color=LIU_C, fontsize=5.5, transform=ax.transAxes)

    # legend
    for x_leg, color, name in [(0.74, DJ_C, 'DJURAEV'), (0.88, LIU_C, 'LIU')]:
        ax.add_patch(FancyBboxPatch((x_leg, 0.93), 0.018, 0.055,
            boxstyle='round,pad=0.002', facecolor=color, alpha=0.88,
            transform=ax.transAxes))
        ax.text(x_leg+0.022, 0.957, name, ha='left', va='center',
                color=color, fontsize=5, transform=ax.transAxes)
    return ax

# ═══════════════════════════════════════════════════════════════════════════════
# IMAGE 1 — HOOK: H2H + Sinclair equivalence
# ═══════════════════════════════════════════════════════════════════════════════
fig1 = plt.figure(figsize=(INCH, INCH), dpi=DPI, facecolor=BG)
gs1  = GridSpec(4, 2, figure=fig1,
                left=0.05, right=0.95, top=0.975, bottom=0.02,
                hspace=0.38, wspace=0.26,
                height_ratios=[0.09, 0.29, 0.33, 0.29])

ax1h = fig1.add_subplot(gs1[0, :])
header(ax1h, subtitle='IWF Career Statistical Analysis  ·  OpenWeightlifting')
nameplate(ax1h, 1)

# ── H2H record ──
ax1r = fig1.add_subplot(gs1[1, :])
card_bg(ax1r)
card_title(ax1r, 'HEAD-TO-HEAD  (same event, same weight class)')
ax1r.text(0.50, 0.62, '2  –  0', ha='center', va='center',
          color=LIU_C, fontsize=38, fontweight='black', transform=ax1r.transAxes)
ax1r.text(0.50, 0.18, 'LIU leads all direct meetings',
          ha='center', va='center', color=LGRAY, fontsize=6.5, transform=ax1r.transAxes)

# ── Two H2H matchups ──
ax1m1 = fig1.add_subplot(gs1[2, 0])
ax1m2 = fig1.add_subplot(gs1[2, 1])

for ax, ev, cat, dt_str, lt, dt_, ls, ds, lbw, dbw in [
    (ax1m1, 'Asian Games 2023', '109 kg Men', '2023-09-30',
     418, 417, 477.8, 462.5, 100.80, 109.00),
    (ax1m2, 'Paris Olympics 2024', '102 kg Men', '2024-08-07',
     406, 404, 462.3, 459.6, 101.75, 102.00),
]:
    card_bg(ax)
    ax.text(0.5, 0.94, ev, ha='center', va='top', color=WHITE,
            fontsize=6.5, fontweight='bold', transform=ax.transAxes)
    ax.text(0.5, 0.80, cat, ha='center', va='top', color=LGRAY,
            fontsize=5, transform=ax.transAxes)

    # LIU column
    ax.text(0.15, 0.60, str(lt), ha='center', va='center', color=LIU_C,
            fontsize=20, fontweight='black', transform=ax.transAxes)
    ax.text(0.15, 0.40, f'Sinclair\n{ls:.1f}', ha='center', va='center',
            color=LIU_C, fontsize=6, alpha=0.8, linespacing=1.5,
            transform=ax.transAxes)
    ax.text(0.15, 0.22, f'BW {lbw} kg', ha='center', va='center',
            color=LGRAY, fontsize=4.8, transform=ax.transAxes)

    # delta
    ax.text(0.50, 0.60, f'+{lt-dt_} kg', ha='center', va='center',
            color=GREEN, fontsize=7, fontweight='bold', transform=ax.transAxes)
    ax.text(0.50, 0.44, 'LIU', ha='center', va='center',
            color=LIU_C, fontsize=5.5, fontweight='bold', transform=ax.transAxes)
    ax.text(0.50, 0.33, 'wins', ha='center', va='center',
            color=LIU_C, fontsize=5, transform=ax.transAxes)

    # DJ column
    ax.text(0.85, 0.60, str(dt_), ha='center', va='center', color=LGRAY,
            fontsize=20, fontweight='black', transform=ax.transAxes)
    ax.text(0.85, 0.40, f'Sinclair\n{ds:.1f}', ha='center', va='center',
            color=LGRAY, fontsize=6, alpha=0.8, linespacing=1.5,
            transform=ax.transAxes)
    ax.text(0.85, 0.22, f'BW {dbw} kg', ha='center', va='center',
            color=LGRAY, fontsize=4.8, transform=ax.transAxes)

    ax.text(0.5, 0.06, dt_str, ha='center', va='bottom', color=DGRAY,
            fontsize=4.5, transform=ax.transAxes)

# ── Equivalence banner ──
ax1e = fig1.add_subplot(gs1[3, :])
card_bg(ax1e)

# big "="
ax1e.text(0.50, 0.72, 'THE SINCLAIR EQUALISER', ha='center', va='center',
          color=LGRAY, fontsize=7, fontweight='bold',
          transform=ax1e.transAxes)

ax1e.text(0.18, 0.42, '477.8', ha='center', va='center', color=LIU_C,
          fontsize=22, fontweight='black', transform=ax1e.transAxes)
ax1e.text(0.18, 0.22, '418 kg total', ha='center', va='center',
          color=LIU_C, fontsize=6, alpha=0.75, transform=ax1e.transAxes)
ax1e.text(0.18, 0.10, '@ 100.8 kg BW', ha='center', va='center',
          color=LIU_C, fontsize=5, alpha=0.55, transform=ax1e.transAxes)

ax1e.text(0.50, 0.38, '=', ha='center', va='center', color=WHITE,
          fontsize=24, fontweight='black', transform=ax1e.transAxes)
ax1e.text(0.50, 0.10, 'Sinclair score', ha='center', va='center',
          color=LGRAY, fontsize=5, transform=ax1e.transAxes)

ax1e.text(0.82, 0.42, '477.3', ha='center', va='center', color=DJ_C,
          fontsize=22, fontweight='black', transform=ax1e.transAxes)
ax1e.text(0.82, 0.22, '446 kg total', ha='center', va='center',
          color=DJ_C, fontsize=6, alpha=0.75, transform=ax1e.transAxes)
ax1e.text(0.82, 0.10, '@ 121.6 kg BW  (+109 kg)', ha='center', va='center',
          color=DJ_C, fontsize=5, alpha=0.55, transform=ax1e.transAxes)

save(fig1, '/home/user/OpenWeightlifting/liu_akbar_1_hook.png')

# ═══════════════════════════════════════════════════════════════════════════════
# IMAGE 2 — CHARTS: Total & Sinclair career progressions
# ═══════════════════════════════════════════════════════════════════════════════
fig2 = plt.figure(figsize=(INCH, INCH), dpi=DPI, facecolor=BG)
gs2  = GridSpec(3, 1, figure=fig2,
                left=0.08, right=0.97, top=0.975, bottom=0.02,
                hspace=0.38,
                height_ratios=[0.07, 0.455, 0.455])

ax2h = fig2.add_subplot(gs2[0, :])
header(ax2h, subtitle='Career Progression  ·  IWF Competitions')
nameplate(ax2h, 2)

# ── Total chart ──
ax2t = fig2.add_subplot(gs2[1])

def annot_total(ax):
    for hd in [datetime(2023,9,30), datetime(2024,8,7)]:
        ax.axvline(hd, color=WHITE, lw=0.7, ls=':', alpha=0.20)
    ax.annotate('', xy=(datetime(2023,9,30), 420), xytext=(datetime(2023,9,30), 427),
                arrowprops=dict(arrowstyle='->', color=LIU_C, lw=1.0))
    ax.annotate('', xy=(datetime(2023,9,30), 415), xytext=(datetime(2023,9,30), 408),
                arrowprops=dict(arrowstyle='->', color=DJ_C,  lw=1.0))
    ax.text(datetime(2023,9,30), 429.5, 'Asian Games\n418 v 417',
            ha='center', va='bottom', color=WHITE, fontsize=4.8, alpha=0.9,
            linespacing=1.5)
    ax.annotate('', xy=(datetime(2024,8,7), 407.5), xytext=(datetime(2024,8,7), 414),
                arrowprops=dict(arrowstyle='->', color=LIU_C, lw=1.0))
    ax.annotate('', xy=(datetime(2024,8,7), 402), xytext=(datetime(2024,8,7), 395),
                arrowprops=dict(arrowstyle='->', color=DJ_C,  lw=1.0))
    ax.text(datetime(2024,8,7), 416, 'Paris Olympics\n406 v 404',
            ha='center', va='bottom', color=WHITE, fontsize=4.8, alpha=0.9,
            linespacing=1.5)
    ax.text(datetime(2022,8,11)+timedelta(days=35), 448,
            '* +109 kg class  (BW 121 kg)', color=DJ_C, fontsize=4.5, alpha=0.65)
    ax.text(0.01, 0.04,
            f'Liu:  {liu_ts*365:+.1f} kg/yr  R²={liu_tr2:.2f}    '
            f'Djuraev:  {dj_ts*365:+.1f} kg/yr  R²={dj_tr2:.2f}',
            transform=ax.transAxes, color=LGRAY, fontsize=5.2, va='bottom')
    ax.set_title('Raw Total (kg)', color=LGRAY, fontsize=7, pad=4)
    p1 = mpatches.Patch(color=LIU_C, label='LIU Huanhua')
    p2 = mpatches.Patch(color=DJ_C,  label='DJURAEV Akbar')
    ax.legend(handles=[p1, p2], loc='upper left', facecolor=CARD,
              edgecolor=BORDER, labelcolor=WHITE, fontsize=6,
              framealpha=0.88, handlelength=1)

scatter_ax(ax2t, liu_dates, liu_tot, dj_dates, dj_tot, 'Total (kg)',
           liu_ts, liu_tb, liu_tx, dj_ts, dj_tb, dj_tx, annot_total)

# ── Sinclair chart ──
ax2s = fig2.add_subplot(gs2[2])

def annot_sinc(ax):
    ax.annotate('', xy=(datetime(2023,9,30), 478.5), xytext=(datetime(2023,9,30), 485),
                arrowprops=dict(arrowstyle='->', color=LIU_C, lw=1.0))
    ax.text(datetime(2023,9,30), 486.5, '477.8',
            ha='center', va='bottom', color=LIU_C, fontsize=5, fontweight='bold')
    ax.annotate('', xy=(datetime(2021,12,7), 481), xytext=(datetime(2021,12,7), 487),
                arrowprops=dict(arrowstyle='->', color=DJ_C,  lw=1.0))
    ax.text(datetime(2021,12,7), 488.5, '480.4',
            ha='center', va='bottom', color=DJ_C, fontsize=5, fontweight='bold')
    ax.text(datetime(2022,7,1), 422,
            'DJ 446 kg @ +109\n= 477.3 Sinclair',
            ha='center', va='center', color=DJ_C, fontsize=4, alpha=0.70,
            linespacing=1.4)
    ax.text(0.01, 0.04,
            f'Liu:  {liu_ss*365:+.1f} pts/yr  R²={liu_sr2:.2f}    '
            f'Djuraev:  {dj_ss*365:+.1f} pts/yr  R²={dj_sr2:.2f}',
            transform=ax.transAxes, color=LGRAY, fontsize=5.2, va='bottom')
    ax.set_title('Sinclair Score  (bodyweight-adjusted)', color=LGRAY, fontsize=7, pad=4)

scatter_ax(ax2s, liu_dates, liu_sin, dj_dates, dj_sin, 'Sinclair',
           liu_ss, liu_sb, liu_sx, dj_ss, dj_sb, dj_sx, annot_sinc)

save(fig2, '/home/user/OpenWeightlifting/liu_akbar_2_charts.png')

# ═══════════════════════════════════════════════════════════════════════════════
# IMAGE 3 — STATS: Sinclair deep-dive + attempt rates
# ═══════════════════════════════════════════════════════════════════════════════
fig3 = plt.figure(figsize=(INCH, INCH), dpi=DPI, facecolor=BG)
gs3  = GridSpec(3, 2, figure=fig3,
                left=0.05, right=0.97, top=0.975, bottom=0.02,
                hspace=0.34, wspace=0.24,
                height_ratios=[0.07, 0.40, 0.53])

ax3h = fig3.add_subplot(gs3[0, :])
header(ax3h, subtitle='Statistical Summary  ·  IWF Data')
nameplate(ax3h, 3)

# ── Sinclair stats card ──
a = fig3.add_subplot(gs3[1, 0])
card_bg(a); card_title(a, 'SINCLAIR  (bodyweight-adjusted)')
row3(a, 0.79, f'{max(liu_sin):.1f}', 'Career Best Sinclair', f'{max(dj_sin):.1f}',
     lc=LGRAY, rc=DJ_C, rb=True)
row3(a, 0.63, f'{np.mean(liu_sin):.1f}', 'Career Mean', f'{np.mean(dj_sin):.1f}',
     lc=LIU_C, rc=LGRAY, lb=True)
row3(a, 0.47,
     f'{np.mean(liu_102_sin):.1f}', '102 kg Mean', f'{np.mean(dj_102_sin):.1f}',
     lc=LIU_C, rc=LGRAY, lb=True)
row3(a, 0.32,
     f'+/-{np.std(liu_102_sin,ddof=1):.1f}', '102 kg Std Dev', f'+/-{np.std(dj_102_sin,ddof=1):.1f}',
     lc=LIU_C, rc=LGRAY, lb=True)
row3(a, 0.17,
     f'{np.std(liu_102_sin,ddof=1)/np.mean(liu_102_sin)*100:.1f}%', '102 kg  CV%',
     f'{np.std(dj_102_sin,ddof=1)/np.mean(dj_102_sin)*100:.1f}%',
     lc=LIU_C, rc=LGRAY, lb=True)
a.text(0.5, 0.03,
       f"Cohen's d = {d_102s:.2f}  |  p = {p_102s:.2f}  (n=5 vs n=3)",
       ha='center', va='bottom', color=DGRAY, fontsize=4.5, transform=a.transAxes)

# ── Career bests (total vs Sinclair) ──
b = fig3.add_subplot(gs3[1, 1])
card_bg(b); card_title(b, 'CAREER BESTS  (Total vs Sinclair)')

headers = ['', 'Total (kg)', 'Sinclair']
col_x   = [0.08, 0.45, 0.82]
row_ys  = [0.79, 0.62, 0.45, 0.28, 0.13]
col_labels = ['', 'Total', 'Sinclair']

for xc, lbl in zip(col_x[1:], col_labels[1:]):
    b.text(xc, 0.90, lbl, ha='center', va='center', color=LGRAY,
           fontsize=5.2, transform=b.transAxes)

rows = [
    ('Best Snatch', '186 kg', '—'),
    ('Best C&J',    '233 kg', '—'),
    ('Best Total',
     f'Liu {max(liu_tot):.0f}  /  DJ {max(dj_tot):.0f}',
     f'Liu {max(liu_sin):.0f}  /  DJ {max(dj_sin):.0f}'),
    ('Career Mean',
     f'Liu {np.mean(liu_tot):.0f}  /  DJ {np.mean(dj_tot):.0f}',
     f'Liu {np.mean(liu_sin):.0f}  /  DJ {np.mean(dj_sin):.0f}'),
]
for (lbl, tv, sv), ry in zip(rows, row_ys):
    b.text(0.02, ry, lbl, ha='left', va='center', color=LGRAY,
           fontsize=5.2, transform=b.transAxes)
    b.text(0.55, ry, tv, ha='center', va='center', color=WHITE,
           fontsize=5.2, transform=b.transAxes)
    b.text(0.92, ry, sv, ha='right', va='center',
           color=LIU_C if 'Liu' in sv and 'DJ' in sv else LGRAY,
           fontsize=5.2, transform=b.transAxes)

b.text(0.5, 0.03,
       'Sinclair neutralises the 28 kg bodyweight advantage',
       ha='center', va='bottom', color=LIU_C, fontsize=4.5,
       fontweight='bold', transform=b.transAxes)

# ── Attempt rates (full width, bottom) ──
bars_section(fig3, gs3[2, :], top=0.97)

save(fig3, '/home/user/OpenWeightlifting/liu_akbar_3_stats.png')

# ═══════════════════════════════════════════════════════════════════════════════
# IMAGE 4 — VERDICT: Who is the better lifter?
# ═══════════════════════════════════════════════════════════════════════════════
liu_cv102        = np.std(liu_102_sin, ddof=1) / np.mean(liu_102_sin) * 100
dj_cv102         = np.std(dj_102_sin,  ddof=1) / np.mean(dj_102_sin)  * 100
liu_overall_rate = np.mean(liu_rates) * 100
dj_overall_rate  = np.mean(dj_rates)  * 100

scorecard = [
    ('Head-to-Head',          '2 – 0',                       '0 – 2',                       'liu'),
    ('Career Best Total',     f'{max(liu_tot):.0f} kg',      f'{max(dj_tot):.0f} kg',        'dj'),
    ('Career Best Sinclair',  f'{max(liu_sin):.1f}',         f'{max(dj_sin):.1f}',           'dj'),
    ('Career Mean Sinclair',  f'{np.mean(liu_sin):.1f}',     f'{np.mean(dj_sin):.1f}',       'liu'),
    ('102 kg Sinclair',       f'{np.mean(liu_102_sin):.1f}', f'{np.mean(dj_102_sin):.1f}',   'liu'),
    ('Consistency  (CV %)',   f'{liu_cv102:.1f}%',           f'{dj_cv102:.1f}%',             'liu'),
    ('Trajectory  (pts/yr)',  f'{liu_ss*365:+.1f}',          f'{dj_ss*365:+.1f}',            'dj'),
    ('Attempt Rate',          f'{liu_overall_rate:.0f}%',    f'{dj_overall_rate:.0f}%',
     'liu' if liu_overall_rate >= dj_overall_rate else 'dj'),
]

liu_score = sum(1 for *_, w in scorecard if w == 'liu')
dj_score  = sum(1 for *_, w in scorecard if w == 'dj')

fig4 = plt.figure(figsize=(INCH, INCH), dpi=DPI, facecolor=BG)
gs4  = GridSpec(3, 1, figure=fig4,
                left=0.05, right=0.97, top=0.975, bottom=0.02,
                hspace=0.22,
                height_ratios=[0.08, 0.55, 0.37])

ax4h = fig4.add_subplot(gs4[0])
header(ax4h, subtitle='Data-driven verdict  ·  Who is the better lifter?')
ax4h.text(0.97, 0.97, '4 / 4', ha='right', va='top', color=DGRAY,
          fontsize=6, transform=ax4h.transAxes)

# ── Scorecard ──
ax4s = fig4.add_subplot(gs4[1])
card_bg(ax4s)
ax4s.text(0.5, 0.97, 'METRIC SCORECARD', ha='center', va='top', color=WHITE,
          fontsize=7.5, fontweight='bold', transform=ax4s.transAxes)

ax4s.text(0.38, 0.89, str(liu_score), ha='center', va='center', color=LIU_C,
          fontsize=20, fontweight='black', transform=ax4s.transAxes)
ax4s.text(0.50, 0.89, '–', ha='center', va='center', color=LGRAY,
          fontsize=14, transform=ax4s.transAxes)
ax4s.text(0.62, 0.89, str(dj_score), ha='center', va='center', color=DJ_C,
          fontsize=20, fontweight='black', transform=ax4s.transAxes)
ax4s.text(0.38, 0.79, 'LIU', ha='center', va='center', color=LIU_C,
          fontsize=5, alpha=0.7, transform=ax4s.transAxes)
ax4s.text(0.62, 0.79, 'DJURAEV', ha='center', va='center', color=DJ_C,
          fontsize=5, alpha=0.7, transform=ax4s.transAxes)

ax4s.plot([0.03, 0.97], [0.745, 0.745], color=BORDER, lw=0.8,
          transform=ax4s.transAxes)

n_rows = len(scorecard)
y0     = 0.700
dy     = y0 / (n_rows + 0.5)

for i, (label, lv, dv, winner) in enumerate(scorecard):
    y = y0 - i * dy - dy * 0.4

    ax4s.add_patch(FancyBboxPatch(
        (0.03, y - dy * 0.44), 0.94, dy * 0.86,
        boxstyle='round,pad=0.003',
        facecolor=LIU_C if winner == 'liu' else DJ_C,
        alpha=0.07, transform=ax4s.transAxes, zorder=1))

    lc = LIU_C if winner == 'liu' else DGRAY
    dc = DJ_C  if winner == 'dj'  else DGRAY
    lw = 'bold' if winner == 'liu' else 'normal'
    dw = 'bold' if winner == 'dj'  else 'normal'

    ax4s.text(0.31, y, lv, ha='right', va='center', color=lc,
              fontsize=6.5, fontweight=lw, transform=ax4s.transAxes, zorder=2)
    ax4s.text(0.50, y, label, ha='center', va='center', color=LGRAY,
              fontsize=5.0, transform=ax4s.transAxes, zorder=2)
    ax4s.text(0.69, y, dv, ha='left', va='center', color=dc,
              fontsize=6.5, fontweight=dw, transform=ax4s.transAxes, zorder=2)

    dot_x = 0.09 if winner == 'liu' else 0.91
    ax4s.plot(dot_x, y, 'o', color=LIU_C if winner == 'liu' else DJ_C,
              ms=3.5, transform=ax4s.transAxes, zorder=3, clip_on=False)

# ── Verdict ──
ax4v = fig4.add_subplot(gs4[2])
card_bg(ax4v)
ax4v.text(0.50, 0.96, 'THE VERDICT', ha='center', va='top', color=LGRAY,
          fontsize=7, fontweight='bold', transform=ax4v.transAxes)
ax4v.plot([0.05, 0.95], [0.87, 0.87], color=BORDER, lw=0.6,
          transform=ax4v.transAxes)
ax4v.text(0.50, 0.76, 'LIU Huanhua', ha='center', va='center', color=LIU_C,
          fontsize=22, fontweight='black', transform=ax4v.transAxes)
ax4v.text(0.50, 0.58, 'POUND-FOR-POUND CHAMPION', ha='center', va='center',
          color=LIU_C, fontsize=8, fontweight='bold', alpha=0.80,
          transform=ax4v.transAxes)
ax4v.plot([0.05, 0.95], [0.50, 0.50], color=BORDER, lw=0.6,
          transform=ax4v.transAxes)

for ri, txt in enumerate([
    '● 2–0 head-to-head in direct competition',
    f"● Higher Sinclair at 102 kg  (Cohen's d = {d_102s:.2f})",
    f'● More consistent: CV {liu_cv102:.1f}% vs {dj_cv102:.1f}% (Djuraev)',
]):
    ax4v.text(0.50, 0.41 - ri * 0.13, txt, ha='center', va='center',
              color=LGRAY, fontsize=5.5, transform=ax4v.transAxes)

ax4v.text(0.50, 0.04,
          f"Djuraev's bigger total (+{int(max(dj_tot) - max(liu_tot))} kg) "
          'reflects bodyweight advantage  ·  OpenWeightlifting',
          ha='center', va='bottom', color=DGRAY, fontsize=4.2,
          transform=ax4v.transAxes)

save(fig4, '/home/user/OpenWeightlifting/liu_akbar_4_verdict.png')

# ═══════════════════════════════════════════════════════════════════════════════
# IMAGE 5 — DISCIPLINES: Snatch & Clean and Jerk career progressions
# ═══════════════════════════════════════════════════════════════════════════════
def best_s(r): return max((a for a in r[4] if a > 0), default=0)
def best_c(r): return max((a for a in r[5] if a > 0), default=0)

liu_sn = np.array([best_s(r) for r in LIU_RAW], dtype=float)
liu_cj = np.array([best_c(r) for r in LIU_RAW], dtype=float)
dj_sn  = np.array([best_s(r) for r in DJ_RAW],  dtype=float)
dj_cj  = np.array([best_c(r) for r in DJ_RAW],  dtype=float)

liu_sns, liu_snb, liu_snr2, liu_snx = linreg(liu_dates, liu_sn)
dj_sns,  dj_snb,  dj_snr2,  dj_snx  = linreg(dj_dates,  dj_sn)
liu_cjs, liu_cjb, liu_cjr2, liu_cjx = linreg(liu_dates, liu_cj)
dj_cjs,  dj_cjb,  dj_cjr2,  dj_cjx  = linreg(dj_dates,  dj_cj)

fig5 = plt.figure(figsize=(INCH, INCH), dpi=DPI, facecolor=BG)
gs5  = GridSpec(3, 1, figure=fig5,
                left=0.08, right=0.97, top=0.975, bottom=0.02,
                hspace=0.40,
                height_ratios=[0.07, 0.455, 0.455])

ax5h = fig5.add_subplot(gs5[0])
header(ax5h, subtitle='Snatch & Clean and Jerk  ·  Career Progression')
ax5h.text(0.97, 0.97, '5 / 5', ha='right', va='top', color=DGRAY,
          fontsize=6, transform=ax5h.transAxes)

# ── Snatch chart ──
ax5s = fig5.add_subplot(gs5[1])

def annot_snatch(ax):
    li = int(np.argmax(liu_sn))
    di = int(np.argmax(dj_sn))
    for idx, dates, vals, color, note in [
        (li, liu_dates, liu_sn, LIU_C, ''),
        (di, dj_dates,  dj_sn,  DJ_C,  '* +109 kg class'),
    ]:
        ax.annotate('', xy=(dates[idx], vals[idx]),
                    xytext=(dates[idx], vals[idx] + 4),
                    arrowprops=dict(arrowstyle='->', color=color, lw=1.0))
        ax.text(dates[idx], vals[idx] + 5.5,
                f'{int(vals[idx])} kg',
                ha='center', va='bottom', color=color, fontsize=5, fontweight='bold')
        if note:
            ax.text(dates[idx], vals[idx] - 8, note,
                    ha='center', color=color, fontsize=4, alpha=0.65)
    ax.set_title('Best Snatch per Competition (kg)', color=LGRAY, fontsize=7, pad=4)
    p1 = mpatches.Patch(color=LIU_C,
                         label=f'LIU  best {int(max(liu_sn))} kg  |  mean {np.mean(liu_sn):.0f} kg')
    p2 = mpatches.Patch(color=DJ_C,
                         label=f'DJURAEV  best {int(max(dj_sn))} kg  |  mean {np.mean(dj_sn):.0f} kg')
    ax.legend(handles=[p1, p2], loc='upper left', facecolor=CARD,
              edgecolor=BORDER, labelcolor=WHITE, fontsize=5.5,
              framealpha=0.88, handlelength=1)
    ax.text(0.01, 0.04,
            f'Liu:  {liu_sns*365:+.1f} kg/yr  R²={liu_snr2:.2f}    '
            f'Djuraev:  {dj_sns*365:+.1f} kg/yr  R²={dj_snr2:.2f}',
            transform=ax.transAxes, color=LGRAY, fontsize=5.2, va='bottom')

scatter_ax(ax5s, liu_dates, liu_sn, dj_dates, dj_sn, 'Snatch (kg)',
           liu_sns, liu_snb, liu_snx, dj_sns, dj_snb, dj_snx, annot_snatch)

# ── C&J chart ──
ax5c = fig5.add_subplot(gs5[2])

def annot_cj(ax):
    li = int(np.argmax(liu_cj))
    di = int(np.argmax(dj_cj))
    for idx, dates, vals, color, note in [
        (li, liu_dates, liu_cj, LIU_C, ''),
        (di, dj_dates,  dj_cj,  DJ_C,  '* +109 kg class'),
    ]:
        ax.annotate('', xy=(dates[idx], vals[idx]),
                    xytext=(dates[idx], vals[idx] + 4),
                    arrowprops=dict(arrowstyle='->', color=color, lw=1.0))
        ax.text(dates[idx], vals[idx] + 5.5,
                f'{int(vals[idx])} kg',
                ha='center', va='bottom', color=color, fontsize=5, fontweight='bold')
        if note:
            ax.text(dates[idx], vals[idx] - 9, note,
                    ha='center', color=color, fontsize=4, alpha=0.65)
    ax.set_title('Best Clean & Jerk per Competition (kg)', color=LGRAY, fontsize=7, pad=4)
    p1 = mpatches.Patch(color=LIU_C,
                         label=f'LIU  best {int(max(liu_cj))} kg  |  mean {np.mean(liu_cj):.0f} kg')
    p2 = mpatches.Patch(color=DJ_C,
                         label=f'DJURAEV  best {int(max(dj_cj))} kg  |  mean {np.mean(dj_cj):.0f} kg')
    ax.legend(handles=[p1, p2], loc='upper left', facecolor=CARD,
              edgecolor=BORDER, labelcolor=WHITE, fontsize=5.5,
              framealpha=0.88, handlelength=1)
    ax.text(0.01, 0.04,
            f'Liu:  {liu_cjs*365:+.1f} kg/yr  R²={liu_cjr2:.2f}    '
            f'Djuraev:  {dj_cjs*365:+.1f} kg/yr  R²={dj_cjr2:.2f}',
            transform=ax.transAxes, color=LGRAY, fontsize=5.2, va='bottom')

scatter_ax(ax5c, liu_dates, liu_cj, dj_dates, dj_cj, 'C&J (kg)',
           liu_cjs, liu_cjb, liu_cjx, dj_cjs, dj_cjb, dj_cjx, annot_cj)

save(fig5, '/home/user/OpenWeightlifting/liu_akbar_5_disciplines.png')
