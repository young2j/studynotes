const Visibility = () => {
  return (
    <>
      <table>
        <thead>
          <tr>
            <th>Invoice #</th>
            <th>Client</th>
            <th>Amount</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>#100</td>
            <td>Pendant Publishing</td>
            <td>$2,000.00</td>
          </tr>
          <tr className="collapse">
            <td>#101</td>
            <td>Kruger Industrial Smoothing</td>
            <td>$545.00</td>
          </tr>
          <tr>
            <td>#102</td>
            <td>J. Peterman</td>
            <td>$10,000.25</td>
          </tr>
        </tbody>
      </table>
    </>
  );
};

export default Visibility;
