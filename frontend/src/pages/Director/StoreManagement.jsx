import { useState } from 'react';
import api from '../../api/axiosClient';

const StoreManagement = () => {
    const [formData, setFormData] = useState({ area: '', counter_count: '', store_type: 'Shop' });
    const [status, setStatus] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await api.post('/reports/stores', {
                area: parseFloat(formData.area),
                counter_count: parseInt(formData.counter_count),
                store_type: formData.store_type
            });
            setStatus(`Успех! Новая точка создана под ID: ${response.data.store_id}`);
            setFormData({ area: '', counter_count: '', store_type: 'Shop' });
        } catch (err) {
            setStatus('Ошибка при создании торговой точки');
        }
    };

    return (
        <div style={{ maxWidth: '400px' }}>
            <h3>Регистрация новой торговой точки</h3>
            {status && <p style={{ color: 'blue' }}>{status}</p>}
            <form onSubmit={handleSubmit} style={{ display: 'flex', flexDirection: 'column', gap: '15px' }}>
                <input type="number" placeholder="Площадь (кв.м)" value={formData.area} onChange={e => setFormData({...formData, area: e.target.value})} required />
                <input type="number" placeholder="Кол-во прилавков" value={formData.counter_count} onChange={e => setFormData({...formData, counter_count: e.target.value})} required />
                <select value={formData.store_type} onChange={e => setFormData({...formData, store_type: e.target.value})}>
                    <option value="Department_Store">Универмаг</option>
                    <option value="Shop">Магазин</option>
                    <option value="Kiosk">Киоск</option>
                    <option value="Tray">Лоток</option>
                </select>
                <button type="submit" style={{ background: '#1890ff', color: 'white', border: 'none', padding: '10px', cursor: 'pointer' }}>Добавить в систему</button>
            </form>
        </div>
    );
};

export default StoreManagement;