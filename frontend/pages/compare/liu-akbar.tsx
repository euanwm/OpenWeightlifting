import HeaderBar from '@/components/molecules/head'
import {
  Table,
  TableHeader,
  TableColumn,
  TableCell,
  TableRow,
  TableBody,
  Chip,
} from '@nextui-org/react'
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
} from 'chart.js'
import { Line } from 'react-chartjs-2'

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
)

type Lift = {
  date: string
  event: string
  category: string
  bw: number
  s1: number
  s2: number
  s3: number
  cj1: number
  cj2: number
  cj3: number
  bestSnatch: number
  bestCJ: number
  total: number
}

const LIU_DATA: Lift[] = [
  { date: '2022-12-05', event: '2022 IWF World Championships', category: '89 kg Men', bw: 88.50, s1: 160, s2: 166, s3: -171, cj1: 205, cj2: 211, cj3: 215, bestSnatch: 166, bestCJ: 215, total: 381 },
  { date: '2023-05-05', event: '2023 Asian Championships', category: '96 kg Men', bw: 89.43, s1: -170, s2: 170, s3: 175, cj1: 210, cj2: -223, cj3: -223, bestSnatch: 175, bestCJ: 210, total: 385 },
  { date: '2023-09-04', event: '2023 IWF World Championships', category: '102 kg Men', bw: 98.52, s1: 171, s2: 176, s3: 180, cj1: -215, cj2: 221, cj3: 224, bestSnatch: 180, bestCJ: 224, total: 404 },
  { date: '2023-09-30', event: '19th Asian Games', category: '109 kg Men', bw: 100.80, s1: -175, s2: 180, s3: 185, cj1: 215, cj2: 227, cj3: 233, bestSnatch: 185, bestCJ: 233, total: 418 },
  { date: '2023-12-04', event: '2023 IWF Grand Prix II', category: '102 kg Men', bw: 100.18, s1: 170, s2: -176, s3: 176, cj1: 210, cj2: 222, cj3: -225, bestSnatch: 176, bestCJ: 222, total: 398 },
  { date: '2024-03-31', event: 'IWF World Cup - Paris 2024 Qual. Event', category: '102 kg Men', bw: 101.84, s1: 175, s2: 181, s3: -186, cj1: 220, cj2: 225, cj3: 232, bestSnatch: 181, bestCJ: 232, total: 413 },
  { date: '2024-08-07', event: 'XXXIII Olympic Games', category: '102 kg Men', bw: 101.75, s1: 178, s2: 183, s3: 186, cj1: 220, cj2: -228, cj3: -233, bestSnatch: 186, bestCJ: 220, total: 406 },
  { date: '2025-05-09', event: 'Asian Championships', category: '102 kg Men', bw: 101.49, s1: 171, s2: 180, s3: -183, cj1: 220, cj2: 230, cj3: -234, bestSnatch: 180, bestCJ: 230, total: 410 },
]

const DJURAEV_DATA: Lift[] = [
  { date: '2017-06-15', event: '2017 IWF Junior World Championships', category: '105 kg Men', bw: 100.30, s1: 155, s2: 159, s3: 162, cj1: -185, cj2: 185, cj3: 190, bestSnatch: 162, bestCJ: 190, total: 352 },
  { date: '2017-11-27', event: '2017 IWF World Championships', category: '105 kg Men', bw: 103.24, s1: 164, s2: 169, s3: 174, cj1: 194, cj2: 199, cj3: -203, bestSnatch: 174, bestCJ: 199, total: 373 },
  { date: '2018-04-20', event: '2018 Asian Junior Championships', category: '105 kg Men', bw: 102.20, s1: 160, s2: 166, s3: 170, cj1: 191, cj2: 197, cj3: 202, bestSnatch: 170, bestCJ: 202, total: 372 },
  { date: '2018-07-07', event: '2018 IWF Junior World Championships', category: '105 kg Men', bw: 102.35, s1: 167, s2: -172, s3: -172, cj1: 195, cj2: 202, cj3: -210, bestSnatch: 167, bestCJ: 202, total: 369 },
  { date: '2018-11-01', event: '2018 IWF World Championships', category: '102 kg Men', bw: 101.67, s1: 173, s2: 178, s3: 180, cj1: 200, cj2: 207, cj3: 212, bestSnatch: 180, bestCJ: 212, total: 392 },
  { date: '2018-12-19', event: '5th International Qatar Cup', category: '109 kg Men', bw: 104.20, s1: 173, s2: -178, s3: 179, cj1: 200, cj2: 207, cj3: 213, bestSnatch: 179, bestCJ: 213, total: 392 },
  { date: '2019-04-18', event: 'Asian Championships', category: '109 kg Men', bw: 108.01, s1: 176, s2: 181, s3: 185, cj1: 215, cj2: 219, cj3: 225, bestSnatch: 185, bestCJ: 225, total: 410 },
  { date: '2019-06-01', event: '2019 IWF Junior World Championships', category: '109 kg Men', bw: 108.55, s1: 173, s2: 177, s3: 182, cj1: 212, cj2: 216, cj3: 0, bestSnatch: 182, bestCJ: 216, total: 398 },
  { date: '2019-09-18', event: '2019 IWF World Championships', category: '109 kg Men', bw: 108.60, s1: -183, s2: 184, s3: 188, cj1: 221, cj2: 226, cj3: 229, bestSnatch: 188, bestCJ: 229, total: 417 },
  { date: '2019-12-10', event: 'IWF World Cup', category: '109 kg Men', bw: 108.10, s1: 176, s2: 181, s3: -184, cj1: 215, cj2: 219, cj3: -229, bestSnatch: 181, bestCJ: 219, total: 400 },
  { date: '2019-12-19', event: '6th Qatar International Cup', category: '109 kg Men', bw: 108.70, s1: 175, s2: 181, s3: 185, cj1: 215, cj2: 220, cj3: 0, bestSnatch: 185, bestCJ: 220, total: 405 },
  { date: '2020-02-08', event: '6th International Solidarity Championships', category: '109 kg Men', bw: 108.65, s1: 177, s2: 182, s3: 189, cj1: -221, cj2: 221, cj3: 0, bestSnatch: 189, bestCJ: 221, total: 410 },
  { date: '2021-04-17', event: '2020 Asian Championships', category: '109 kg Men', bw: 108.50, s1: 188, s2: 194, s3: -197, cj1: 225, cj2: 234, cj3: -238, bestSnatch: 194, bestCJ: 234, total: 428 },
  { date: '2021-07-23', event: 'XXXII Olympic Games (Tokyo)', category: '109 kg Men', bw: 109.00, s1: -189, s2: 189, s3: 193, cj1: 227, cj2: -234, cj3: 237, bestSnatch: 193, bestCJ: 237, total: 430 },
  { date: '2021-12-07', event: '2021 IWF World Championships', category: '109 kg Men', bw: 108.90, s1: 187, s2: 192, s3: 195, cj1: 226, cj2: 232, cj3: 238, bestSnatch: 195, bestCJ: 238, total: 433 },
  { date: '2022-08-11', event: '5th Islamic Solidarity Games', category: '+109 kg Men', bw: 121.60, s1: -190, s2: 190, s3: 200, cj1: 231, cj2: 242, cj3: 246, bestSnatch: 200, bestCJ: 246, total: 446 },
  { date: '2023-05-05', event: '2023 Asian Championships', category: '+109 kg Men', bw: 126.76, s1: 189, s2: -195, s3: 195, cj1: 230, cj2: -240, cj3: 242, bestSnatch: 195, bestCJ: 242, total: 437 },
  { date: '2023-09-04', event: '2023 IWF World Championships', category: '109 kg Men', bw: 108.84, s1: 182, s2: -189, s3: 189, cj1: 220, cj2: 226, cj3: -231, bestSnatch: 189, bestCJ: 226, total: 415 },
  { date: '2023-09-30', event: '19th Asian Games', category: '109 kg Men', bw: 109.00, s1: 180, s2: 184, s3: 189, cj1: -220, cj2: 222, cj3: 228, bestSnatch: 189, bestCJ: 228, total: 417 },
  { date: '2024-02-03', event: 'Asian Championships', category: '102 kg Men', bw: 101.70, s1: 175, s2: 180, s3: -183, cj1: 214, cj2: -219, cj3: 220, bestSnatch: 180, bestCJ: 220, total: 400 },
  { date: '2024-03-31', event: 'IWF World Cup - Paris 2024 Qual. Event', category: '109 kg Men', bw: 108.50, s1: 180, s2: 185, s3: 189, cj1: 220, cj2: 227, cj3: 0, bestSnatch: 189, bestCJ: 227, total: 416 },
  { date: '2024-08-07', event: 'XXXIII Olympic Games (Paris)', category: '102 kg Men', bw: 102.00, s1: 180, s2: 185, s3: -189, cj1: 219, cj2: -224, cj3: -232, bestSnatch: 185, bestCJ: 219, total: 404 },
  { date: '2025-05-09', event: 'Asian Championships', category: '109 kg Men', bw: 108.91, s1: 180, s2: 183, s3: -189, cj1: 217, cj2: 223, cj3: 0, bestSnatch: 183, bestCJ: 223, total: 406 },
  { date: '2025-10-02', event: '2025 IWF World Championships', category: '110 kg Men', bw: 109.85, s1: 189, s2: 193, s3: 196, cj1: 227, cj2: 232, cj3: -245, bestSnatch: 196, bestCJ: 232, total: 428 },
]

type AttemptRate = { made: number; total: number; pct: number }

function computeAttemptRates(data: Lift[]) {
  const slots = [
    { key: 's1' as keyof Lift },
    { key: 's2' as keyof Lift },
    { key: 's3' as keyof Lift },
    { key: 'cj1' as keyof Lift },
    { key: 'cj2' as keyof Lift },
    { key: 'cj3' as keyof Lift },
  ]
  return slots.map(({ key }) => {
    let made = 0
    let total = 0
    data.forEach(lift => {
      const v = lift[key] as number
      if (v !== 0) {
        total++
        if (v > 0) made++
      }
    })
    return { made, total, pct: total > 0 ? Math.round((made / total) * 100) : 0 } as AttemptRate
  })
}

const HEAD_TO_HEAD = [
  {
    date: '2023-09-30',
    event: '19th Asian Games',
    category: '109 kg Men',
    liuTotal: 418,
    liuSnatch: 185,
    liuCJ: 233,
    liuBW: 100.80,
    djuraevTotal: 417,
    djuraevSnatch: 189,
    djuraevCJ: 228,
    djuraevBW: 109.00,
    winner: 'liu' as const,
    margin: 1,
  },
  {
    date: '2024-08-07',
    event: 'XXXIII Olympic Games (Paris)',
    category: '102 kg Men',
    liuTotal: 406,
    liuSnatch: 186,
    liuCJ: 220,
    liuBW: 101.75,
    djuraevTotal: 404,
    djuraevSnatch: 185,
    djuraevCJ: 219,
    djuraevBW: 102.00,
    winner: 'liu' as const,
    margin: 2,
  },
]

function formatAttempt(v: number) {
  if (v === 0) return '—'
  if (v < 0) return <span className="text-[#9d3d3d]">{Math.abs(v)}</span>
  return <span className="text-[#598138]">{v}</span>
}

function StatCard({ label, liuVal, djuraevVal }: { label: string; liuVal: string | number; djuraevVal: string | number }) {
  const liuNum = typeof liuVal === 'number' ? liuVal : parseFloat(String(liuVal))
  const djNum = typeof djuraevVal === 'number' ? djuraevVal : parseFloat(String(djuraevVal))
  const liuWins = liuNum > djNum
  const djWins = djNum > liuNum

  return (
    <div className="flex flex-col items-center bg-[#111] rounded-xl p-4 border border-[#222]">
      <span className="text-xs text-[#A0A0A0] mb-3 text-center uppercase tracking-widest">{label}</span>
      <div className="flex items-center gap-4 w-full justify-around">
        <span className={`text-2xl font-bold ${liuWins ? 'text-[#ffce00]' : 'text-white'}`}>{liuVal}</span>
        <span className="text-[#484848] text-sm">vs</span>
        <span className={`text-2xl font-bold ${djWins ? 'text-[#ffce00]' : 'text-white'}`}>{djuraevVal}</span>
      </div>
    </div>
  )
}

function AttemptBar({ rate, color }: { rate: AttemptRate; color: string }) {
  return (
    <div className="flex items-center gap-2">
      <div className="flex-1 bg-[#1a1a1a] rounded-full h-2">
        <div
          className="h-2 rounded-full transition-all"
          style={{ width: `${rate.pct}%`, backgroundColor: color }}
        />
      </div>
      <span className="text-xs text-[#A0A0A0] w-12 text-right">{rate.pct}%</span>
    </div>
  )
}

export default function LiuAkbarComparePage() {
  const liuRates = computeAttemptRates(LIU_DATA)
  const djuraevRates = computeAttemptRates(DJURAEV_DATA)

  const liuBestSnatch = Math.max(...LIU_DATA.map(d => d.bestSnatch))
  const liuBestCJ = Math.max(...LIU_DATA.map(d => d.bestCJ))
  const liuBestTotal = Math.max(...LIU_DATA.map(d => d.total))

  const djBestSnatch = Math.max(...DJURAEV_DATA.map(d => d.bestSnatch))
  const djBestCJ = Math.max(...DJURAEV_DATA.map(d => d.bestCJ))
  const djBestTotal = Math.max(...DJURAEV_DATA.map(d => d.total))

  const chartData = {
    labels: LIU_DATA.map(d => d.date.slice(0, 7)),
    datasets: [
      {
        label: 'LIU Huanhua — Total',
        data: LIU_DATA.map(d => d.total),
        borderColor: '#00B0F0',
        backgroundColor: '#00B0F020',
        tension: 0.3,
        pointRadius: 5,
      },
    ],
  }

  const djChartData = {
    labels: DJURAEV_DATA.map(d => d.date.slice(0, 7)),
    datasets: [
      {
        label: 'DJURAEV Akbar — Total',
        data: DJURAEV_DATA.map(d => d.total),
        borderColor: '#ffce00',
        backgroundColor: '#ffce0020',
        tension: 0.3,
        pointRadius: 5,
      },
    ],
  }

  const mergedDates = Array.from(
    new Set([...LIU_DATA.map(d => d.date), ...DJURAEV_DATA.map(d => d.date)])
  ).sort()

  const liuByDate = Object.fromEntries(LIU_DATA.map(d => [d.date, d.total]))
  const djByDate = Object.fromEntries(DJURAEV_DATA.map(d => [d.date, d.total]))

  const overlapDates = mergedDates.filter(d => liuByDate[d] !== undefined && djByDate[d] !== undefined)

  const overlapChartData = {
    labels: overlapDates.map(d => d.slice(0, 7)),
    datasets: [
      {
        label: 'LIU Huanhua',
        data: overlapDates.map(d => liuByDate[d]),
        borderColor: '#00B0F0',
        backgroundColor: '#00B0F020',
        tension: 0.3,
        pointRadius: 6,
      },
      {
        label: 'DJURAEV Akbar',
        data: overlapDates.map(d => djByDate[d]),
        borderColor: '#ffce00',
        backgroundColor: '#ffce0020',
        tension: 0.3,
        pointRadius: 6,
      },
    ],
  }

  const chartOptions = {
    plugins: { legend: { display: true, labels: { color: '#A0A0A0' } } },
    scales: {
      x: { grid: { display: false }, ticks: { color: '#A0A0A0', font: { size: 11 } } },
      y: { grid: { display: false }, ticks: { color: '#A0A0A0', font: { size: 11 } } },
    },
    aspectRatio: 2.5,
  }

  const attemptLabels = ['1st Snatch', '2nd Snatch', '3rd Snatch', '1st C&J', '2nd C&J', '3rd C&J']

  return (
    <>
      <HeaderBar />
      <div className="max-w-6xl mx-auto px-4 py-6 space-y-10">

        {/* Hero */}
        <div className="flex flex-col sm:flex-row items-center justify-center gap-4 text-center">
          <div className="flex-1">
            <p className="text-xs text-[#A0A0A0] uppercase tracking-widest mb-1">China 🇨🇳</p>
            <h1 className="text-3xl sm:text-4xl font-bold text-[#00B0F0]">LIU Huanhua</h1>
            <p className="text-sm text-[#A0A0A0] mt-1">{LIU_DATA.length} IWF competitions</p>
          </div>
          <div className="text-3xl font-black text-[#484848] px-4">VS</div>
          <div className="flex-1">
            <p className="text-xs text-[#A0A0A0] uppercase tracking-widest mb-1">Uzbekistan 🇺🇿</p>
            <h1 className="text-3xl sm:text-4xl font-bold text-[#ffce00]">DJURAEV Akbar</h1>
            <p className="text-sm text-[#A0A0A0] mt-1">{DJURAEV_DATA.length} IWF competitions</p>
          </div>
        </div>

        {/* Career stats */}
        <section>
          <h2 className="text-lg font-semibold text-[#A0A0A0] uppercase tracking-widest mb-4">Career Bests</h2>
          <div className="grid grid-cols-2 sm:grid-cols-4 gap-3">
            <StatCard label="Best Snatch" liuVal={liuBestSnatch} djuraevVal={djBestSnatch} />
            <StatCard label="Best C&J" liuVal={liuBestCJ} djuraevVal={djBestCJ} />
            <StatCard label="Best Total" liuVal={liuBestTotal} djuraevVal={djBestTotal} />
            <StatCard label="IWF Competitions" liuVal={LIU_DATA.length} djuraevVal={DJURAEV_DATA.length} />
          </div>
          <p className="text-xs text-[#484848] mt-2">
            * Djuraev&apos;s 446kg total and 200kg snatch were set at the 5th Islamic Solidarity Games 2022 competing at +109kg (121.6kg bodyweight).
          </p>
        </section>

        {/* Head to head */}
        <section>
          <h2 className="text-lg font-semibold text-[#A0A0A0] uppercase tracking-widest mb-4">
            Head-to-Head — Same Event, Same Weight Class
          </h2>
          <div className="space-y-4">
            {HEAD_TO_HEAD.map((match, i) => (
              <div key={i} className="bg-[#111] rounded-xl border border-[#222] p-4">
                <div className="flex flex-col sm:flex-row items-center gap-2 mb-3">
                  <span className="text-sm font-semibold text-white">{match.event}</span>
                  <span className="text-xs text-[#484848]">·</span>
                  <span className="text-xs text-[#A0A0A0]">{match.date}</span>
                  <span className="text-xs text-[#484848]">·</span>
                  <Chip size="sm" variant="flat" className="text-xs">{match.category}</Chip>
                </div>
                <div className="grid grid-cols-3 text-center gap-2">
                  <div>
                    <p className={`text-xs uppercase tracking-wider mb-1 ${match.winner === 'liu' ? 'text-[#00B0F0]' : 'text-[#A0A0A0]'}`}>
                      LIU {match.winner === 'liu' && '🏆'}
                    </p>
                    <p className="text-xs text-[#A0A0A0]">BW {match.liuBW} kg</p>
                    <p className="text-lg font-bold text-white mt-1">{match.liuTotal} kg</p>
                    <p className="text-xs text-[#A0A0A0]">{match.liuSnatch} / {match.liuCJ}</p>
                  </div>
                  <div className="flex flex-col items-center justify-center">
                    <span className="text-[#484848] text-xl font-black">|</span>
                    <span className="text-xs text-[#A0A0A0] mt-1">Δ {match.margin} kg</span>
                  </div>
                  <div>
                    <p className={`text-xs uppercase tracking-wider mb-1 ${match.winner === 'djuraev' ? 'text-[#ffce00]' : 'text-[#A0A0A0]'}`}>
                      DJURAEV {match.winner === 'djuraev' && '🏆'}
                    </p>
                    <p className="text-xs text-[#A0A0A0]">BW {match.djuraevBW} kg</p>
                    <p className="text-lg font-bold text-white mt-1">{match.djuraevTotal} kg</p>
                    <p className="text-xs text-[#A0A0A0]">{match.djuraevSnatch} / {match.djuraevCJ}</p>
                  </div>
                </div>
              </div>
            ))}
            <p className="text-xs text-[#484848]">
              Liu leads the direct head-to-head record 2–0, winning both encounters by narrow margins (1 kg and 2 kg).
            </p>
          </div>
        </section>

        {/* Shared-date totals chart */}
        <section>
          <h2 className="text-lg font-semibold text-[#A0A0A0] uppercase tracking-widest mb-4">
            Total on Same Competition Dates
          </h2>
          <Line data={overlapChartData} options={chartOptions} />
        </section>

        {/* Individual career progression */}
        <section>
          <h2 className="text-lg font-semibold text-[#A0A0A0] uppercase tracking-widest mb-4">
            Career Total Progression
          </h2>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <p className="text-sm text-[#00B0F0] mb-2 text-center font-semibold">LIU Huanhua</p>
              <Line data={chartData} options={{ ...chartOptions, aspectRatio: 1.8 }} />
            </div>
            <div>
              <p className="text-sm text-[#ffce00] mb-2 text-center font-semibold">DJURAEV Akbar</p>
              <Line data={djChartData} options={{ ...chartOptions, aspectRatio: 1.8 }} />
            </div>
          </div>
        </section>

        {/* Attempt success rates */}
        <section>
          <h2 className="text-lg font-semibold text-[#A0A0A0] uppercase tracking-widest mb-4">
            Attempt Success Rates
          </h2>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div className="bg-[#111] rounded-xl border border-[#222] p-5 space-y-3">
              <p className="text-sm font-semibold text-[#00B0F0] mb-2">LIU Huanhua</p>
              {liuRates.map((rate, i) => (
                <div key={i}>
                  <div className="flex justify-between text-xs text-[#A0A0A0] mb-1">
                    <span>{attemptLabels[i]}</span>
                    <span>{rate.made}/{rate.total}</span>
                  </div>
                  <AttemptBar rate={rate} color="#00B0F0" />
                </div>
              ))}
            </div>
            <div className="bg-[#111] rounded-xl border border-[#222] p-5 space-y-3">
              <p className="text-sm font-semibold text-[#ffce00] mb-2">DJURAEV Akbar</p>
              {djuraevRates.map((rate, i) => (
                <div key={i}>
                  <div className="flex justify-between text-xs text-[#A0A0A0] mb-1">
                    <span>{attemptLabels[i]}</span>
                    <span>{rate.made}/{rate.total}</span>
                  </div>
                  <AttemptBar rate={rate} color="#ffce00" />
                </div>
              ))}
            </div>
          </div>
        </section>

        {/* Full history tables */}
        <section>
          <h2 className="text-lg font-semibold text-[#A0A0A0] uppercase tracking-widest mb-4">
            LIU Huanhua — Full IWF History
          </h2>
          <div className="overflow-x-auto">
            <Table aria-label="LIU Huanhua competition history">
              <TableHeader>
                <TableColumn>Date</TableColumn>
                <TableColumn>Event</TableColumn>
                <TableColumn>Cat.</TableColumn>
                <TableColumn>BW</TableColumn>
                <TableColumn>S1</TableColumn>
                <TableColumn>S2</TableColumn>
                <TableColumn>S3</TableColumn>
                <TableColumn>CJ1</TableColumn>
                <TableColumn>CJ2</TableColumn>
                <TableColumn>CJ3</TableColumn>
                <TableColumn>Total</TableColumn>
              </TableHeader>
              <TableBody>
                {[...LIU_DATA].reverse().map((lift, i) => (
                  <TableRow key={i}>
                    <TableCell>{lift.date}</TableCell>
                    <TableCell>{lift.event}</TableCell>
                    <TableCell>{lift.category}</TableCell>
                    <TableCell>{lift.bw}</TableCell>
                    <TableCell>{formatAttempt(lift.s1)}</TableCell>
                    <TableCell>{formatAttempt(lift.s2)}</TableCell>
                    <TableCell>{formatAttempt(lift.s3)}</TableCell>
                    <TableCell>{formatAttempt(lift.cj1)}</TableCell>
                    <TableCell>{formatAttempt(lift.cj2)}</TableCell>
                    <TableCell>{formatAttempt(lift.cj3)}</TableCell>
                    <TableCell><span className="font-bold">{lift.total}</span></TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>
        </section>

        <section>
          <h2 className="text-lg font-semibold text-[#A0A0A0] uppercase tracking-widest mb-4">
            DJURAEV Akbar — Full IWF History
          </h2>
          <div className="overflow-x-auto">
            <Table aria-label="DJURAEV Akbar competition history">
              <TableHeader>
                <TableColumn>Date</TableColumn>
                <TableColumn>Event</TableColumn>
                <TableColumn>Cat.</TableColumn>
                <TableColumn>BW</TableColumn>
                <TableColumn>S1</TableColumn>
                <TableColumn>S2</TableColumn>
                <TableColumn>S3</TableColumn>
                <TableColumn>CJ1</TableColumn>
                <TableColumn>CJ2</TableColumn>
                <TableColumn>CJ3</TableColumn>
                <TableColumn>Total</TableColumn>
              </TableHeader>
              <TableBody>
                {[...DJURAEV_DATA].reverse().map((lift, i) => (
                  <TableRow key={i}>
                    <TableCell>{lift.date}</TableCell>
                    <TableCell>{lift.event}</TableCell>
                    <TableCell>{lift.category}</TableCell>
                    <TableCell>{lift.bw}</TableCell>
                    <TableCell>{formatAttempt(lift.s1)}</TableCell>
                    <TableCell>{formatAttempt(lift.s2)}</TableCell>
                    <TableCell>{formatAttempt(lift.s3)}</TableCell>
                    <TableCell>{formatAttempt(lift.cj1)}</TableCell>
                    <TableCell>{formatAttempt(lift.cj2)}</TableCell>
                    <TableCell>{formatAttempt(lift.cj3)}</TableCell>
                    <TableCell><span className="font-bold">{lift.total}</span></TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>
        </section>

        <footer className="text-center text-xs text-[#484848] pb-6">
          Data sourced from IWF event results via OpenWeightlifting. Compiled from event_data/IWF/*.csv
        </footer>
      </div>
    </>
  )
}
