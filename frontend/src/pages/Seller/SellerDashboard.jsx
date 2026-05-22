import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import CreateSale from './CreateSale';
import TransferProduct from './TransferProduct';

const SellerDashboard = () => {
    const [activeTab, setActiveTab] = useState('sale');
    const navigate = useNavigate();

    const handleLogout = () => {
        localStorage.removeItem('token');
        localStorage.removeItem('role');
        navigate('/login');
    };

    return (
        <div style={{ padding: '20px', maxWidth: '800px', margin: '0 auto' }}>
            <header style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', borderBottom: '2px solid #eee', paddingBottom: '10px', marginBottom: '20px' }}>
                <h2>Рабочее место Продавца</h2>
                <button onClick={handleLogout} style={{ background: '#ff4d4f', color: 'white', border: 'none', padding: '8px 16px', borderRadius: '4px', cursor: 'pointer' }}>
                    Выйти
                </button>
            </header>

            <div style={{ marginBottom: '20px', display: 'flex', gap: '10px' }}>
                <button
                    onClick={() => setActiveTab('sale')}
                    style={{ padding: '10px 20px', background: activeTab === 'sale' ? '#1890ff' : '#f0f0f0', color: activeTab === 'sale' ? 'white' : 'black', border: 'none', borderRadius: '4px', cursor: 'pointer' }}
                >
                    Оформление чека
                </button>
                <button
                    onClick={() => setActiveTab('transfer')}
                    style={{ padding: '10px 20px', background: activeTab === 'transfer' ? '#1890ff' : '#f0f0f0', color: activeTab === 'transfer' ? 'white' : 'black', border: 'none', borderRadius: '4px', cursor: 'pointer' }}
                >
                    Перемещение товара
                </button>
            </div>

            <main>
                {activeTab === 'sale' && <CreateSale />}
                {activeTab === 'transfer' && <TransferProduct />}
            </main>
        </div>
    );
};

export default SellerDashboard;