using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using System.Net.Http;
using System.Diagnostics;


namespace ClientUtenteCS.Forms
{
    public partial class FormViewTickets : Form
    {
        private Ticket[] arrayTickets;
        public FormViewTickets()
        {
            InitializeComponent();            
            loadTickets();

        }

        private void FormViewTickets_Load(object sender, EventArgs e)
        {

        }


        private async void loadTickets()
        {
            Richieste getTickets = new Richieste();


            try
            {
                var mail = new Dictionary<string, string>
                {
                    {"email", FormLogin.emailUtente}
                };
                var datiDaInviare = new FormUrlEncodedContent(mail);
                var result = await getTickets.HttpPostAsync("http://localhost:8000/getickets", datiDaInviare);
                // da risolvere: gestire il json in arrivo come un array di ticket
              //  arrayTickets = result;
            }

            catch (Exception exc)
            {
                Debug.WriteLine(exc);
                MessageBox.Show("Non è possibile contattare il server in questo momento. Riprova più tardi.",
                    "Server error", MessageBoxButtons.OK, MessageBoxIcon.Warning);

            }
        }
    }
}
