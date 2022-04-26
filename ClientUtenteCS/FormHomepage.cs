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

namespace ClientUtenteCS
{
    public partial class FormHomepage : Form
    {

        // campi
        private Button bottoneCorrente;
        private Random random;
        private int indiceTemp;
        private Form formAttivo;
        private string nomeUtente;

        //costruttori
        public FormHomepage()
        {
            InitializeComponent();
            random = new Random();
            buttonBackHome.Visible = false;
            this.Text = String.Empty;
            this.ControlBox = false;
            getNomeUtente();
            
        }

        // metodi
        private Color SelezionaTemaColore()
        {
            ColoreTema tema = new ColoreTema();
            int indice = random.Next(tema.ListaColori.Count);
            // per randomizzare il tema del colore della finestra form
            while (indiceTemp == indice)
            {
               indice=  random.Next(tema.ListaColori.Count);
            }
            indiceTemp = indice;
            string colore = tema.ListaColori[indice];
            return ColorTranslator.FromHtml(colore);
        }

        // cambiamenti del colore e del font dei bottoni al click
        private void bottoneAttivo(object bottone)
        {
            if (bottone != null)
            {
                if (bottoneCorrente != (Button)bottone)
                {
                    bottoneDisattivo();
                    Color colore = SelezionaTemaColore();
                    bottoneCorrente = (Button)bottone;
                    bottoneCorrente.BackColor = colore;
                    bottoneCorrente.ForeColor = Color.White;
                    bottoneCorrente.Font = new System.Drawing.Font("Microsoft Sans Serif", 10.00F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(0)));
                    panelTitle.BackColor = colore;
                    buttonBackHome.Visible = true;
                }
            }
        }

        private void bottoneDisattivo ()
        {
            foreach (Control bottoni in panelMenu.Controls)
            {
                if (bottoni.GetType() == typeof(Button))
                {
                    bottoni.BackColor = Color.FromArgb(178, 34, 34);
                    bottoni.Font = new System.Drawing.Font("Microsoft Sans Serif", 8.25F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(0)));
                }
            }
        }

        // metodo per aprire altri form all'interno del pannello dove si cliccano le varie voci del menu
        private void OpenForms(Form formFiglio, object bottone)
        {
            if (formAttivo != null)
            {
                formAttivo.Close();
            }
            bottoneAttivo(bottone);
            formAttivo = formFiglio;
            formFiglio.TopLevel = false;
            formFiglio.FormBorderStyle= FormBorderStyle.None;
            formFiglio.Dock= DockStyle.Fill;
            this.panelSottostante.Controls.Add(formFiglio);
            this.panelSottostante.Tag = formFiglio;
            formFiglio.BringToFront();
            formFiglio.Show();
         //   labelTitle.Text = formFiglio.Text;

        }


        private void buttonAddTicket_Click(object sender, EventArgs e)
        {
            labelTitle.Text = "CREAZIONE TICKET";
            OpenForms(new Forms.FormAddTicket(), sender);
            // implementare qui form per selezionare il professionista
        }

        private void buttonViewTicket_Click(object sender, EventArgs e)
        {
            OpenForms(new Forms.FormViewTickets(), sender);
            labelTitle.Text = "I TUOI TICKET";
        }

        private void buttonPreventivi_Click(object sender, EventArgs e)
        {
            bottoneAttivo(sender);
            labelTitle.Text = "I TUOI PREVENTIVI";
        }

        private void buttonFatture_Click(object sender, EventArgs e)
        {
            bottoneAttivo(sender);
            labelTitle.Text = "LE TUE FATTURE";
        }

        private void buttonReview_Click(object sender, EventArgs e)
        {
            bottoneAttivo(sender);
            labelTitle.Text = "LE TUE OPINIONI";
        }

        private void buttonBackHome_Click(object sender, EventArgs e)
        {
            if (formAttivo != null)
            {
                formAttivo.Close();
            }
            Reset();
        }
        // ritorno all'homepage
        private void Reset()
        {
            bottoneDisattivo();
            labelTitle.Text = "Benvenuto, " + nomeUtente;
            panelTitle.BackColor = Color.FromArgb(128, 0, 0);
            bottoneCorrente = null;
            buttonBackHome.Visible = false;
        }

        private async void getNomeUtente()
        {
            Richieste getNome = new Richieste();

            try
            {
                var mail = new Dictionary<string, string>
                {
                    {"email", FormLogin.emailUtente}
                };
                var datiDaInviare = new FormUrlEncodedContent(mail);
                var result = await getNome.HttpPostAsync("http://localhost:8000/getnome", datiDaInviare);

                if (result.Equals("Errore generico"))
                {
                    nomeUtente = "utente";
                }
                else
                {
                    nomeUtente = result;
                }
                labelTitle.Text = "Benvenuto, " + nomeUtente;
            }
             catch (Exception exc)
            {
                Debug.WriteLine(exc);
            }
        }

        private void buttonExit_Click(object sender, EventArgs e)
        {
            Application.Exit();
        }
    }
}
