import { useState } from 'react';
import api from '../../api/axiosClient';

const ProfitabilityReport = () => {
    const [filters, setFilters] = useState({ store_id: '', start_date: '2026-04-01', end_date: '2026-04-30' });
    const [report, setReport] = useState(null);

    const handleSearch = async (e) => {
        e.preventDefault();
        try {
            const response = await api.get(`/reports/profitability`, { params: filters });
            setReport(response.data);
        } catch (err) {
            alert('Ошибка при получении данных рентабельности');
        }
    };

    return (
        <div>
            <h3>Рентабельность торговой точки</h3>
            <form onSubmit={handleSearch} style={{ display: 'flex', gap: '10px', flexWrap: 'wrap', marginBottom: '20px' }}>
                <input type="number" placeholder="ID Точки" value={filters.store_id} onChange={e => setFilters({...filters, store_id: e.target.value})} required style={{ padding: '5px' }} />
                <input type="date" value={filters.start_date} onChange={e => setFilters({...filters, start_date: e.target.value})} style={{ padding: '5px' }} />
                <input type="date" value={filters.end_date} onChange={e => setFilters({...filters, end_date: e.target.value})} style={{ padding: '5px' }} />
                <button type="submit" style={{ padding: '5px 15px', cursor: 'pointer' }}>Рассчитать</button>
            </form>

            {report && (
                <div style={{ padding: '15px', border: '1px solid #52c41a', borderRadius: '4px', background: '#f6ffed' }}>
                    <p><strong>Общая выручка:</strong> {report.total_revenue.toLocaleString()} ₽</p>
                    <p><strong>Накладные расходы (Аренда+ЗП):</strong> {report.total_overhead.toLocaleString()} ₽</p>
                    <p><strong>Коэффициент рентабельности:</strong>
                        <span style={{ color: report.profitability_ratio > 1 ? 'green' : 'red', fontWeight: 'bold', marginLeft: '10px' }}>
                            {report.profitability_ratio ? report.profitability_ratio.toFixed(2) : 'Н/Д'}
                        </span>
                    </p>
                </div>
            )}
        </div>
    );
};

export default ProfitabilityReport;