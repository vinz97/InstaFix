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
using Newtonsoft.Json;

namespace ClientUtenteCS.Forms
{
    public partial class FormViewTickets : Form
    {
       // private Ticket[] arrayTickets;
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

                // utilizzo della funzione JsonConvert della libreria Newtonsoft scaricabile tramite pacchetti NuGet
                // su VS: progetto/Gestisci pacchetti NuGet (non basta solo l'include)
                List<Ticket> arrayTickets = JsonConvert.DeserializeObject<List<Ticket>>(result);
                //        foreach (var item in arrayTickets) Console.WriteLine(item);
                if (arrayTickets != null)
                {
                    foreach (var item in arrayTickets)
                    {
                        listViewTickets.Items.Add(new ListViewItem(new String[] { item.Id.ToString(), item.Titolo, item.Categoria, item.Stato }));

                    }
                }
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
