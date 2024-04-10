Simulação de Ataques usando MQTT

# Relatório

Este relatório apresenta uma análise detalhada dos fundamentos de segurança da informação aplicados a ambientes de Internet das Coisas (IoT), com foco na Simulação de Ataques MQTT. Exploraremos conceitos essenciais, como a proteção de dados em trânsito e em armazenamento, bem como a autenticação e autorização da comunicação. Além disso, discutiremos a relevância e a eficácia de cada uma dessas estratégias de proteção, destacando sua importância para garantir a segurança e integridade das informações em ambientes IoT.

*Pilares do CIA Triad:*

- *Integridade*: Refere-se à garantia de que os dados permaneçam íntegros, precisos e confiáveis ao longo do tempo e durante todo o processo de manipulação.
  
- *Disponibilidade*: Garante que os recursos de informação estejam disponíveis quando necessários.
  
- *Confidencialidade*: Visa proteger informações sensíveis de serem acessadas por usuários não autorizados.

*Pergunta:* 

O que acontece se você utilizar o mesmo ClientID em outra máquina ou sessão do navegador? Algum pilar do CIA Triad é violado com isso?

*Resposta:* 

Quando uma segunda sessão do navegador é conectada com o mesmo ClientID, a primeira sessão será desconectada automaticamente. Isso resulta em uma violação do pilar de Integridade. Utilizar o mesmo ClientID em diferentes máquinas ou sessões do navegador pode comprometer a integridade dos dados associados a esse identificador, uma vez que não há garantia de que as alterações feitas em uma máquina ou sessão sejam refletidas de forma consistente em todas as outras instâncias.

*Pergunta:* 

Com os parâmetros de resources, algum pilar do CIA Triad pode ser facilmente violado?

*Resposta:* 

Os parâmetros de recursos definidos estão relacionados à alocação de recursos de computação para o contêiner MQTT5, como CPU e memória. Embora esses parâmetros possam influenciar o desempenho e a disponibilidade do serviço, não estão diretamente ligados aos pilares da CIA Triad. No entanto, se as limitações dos parâmetros de recursos forem significativas o suficiente, isso pode afetar os três pilares da CIA Triad.

*Pergunta:* 

Sem autenticação (repare que a variável allow_anonymous está como true), como a parte de confidencialidade pode ser violada?

*Resposta:* 

Sem autenticação adequada, várias violações de confidencialidade podem ocorrer. O acesso não autorizado permite que qualquer pessoa, incluindo indivíduos mal-intencionados, acesse dados confidenciais. A intercepção de dados pode ocorrer, onde os dados transmitidos são interceptados por terceiros. O roubo de identidade é facilitado, pois não há verificação da identidade dos usuários. Além disso, a falta de autenticação aumenta o risco de vazamento de dados sensíveis, incluindo informações privadas.

*Pergunta:* 

Tente simular uma violação do pilar de Confidencialidade.

*Resposta:* 

A demonstração no vídeo mostra uma conexão bem-sucedida com um broker remoto usando um web client, seguida por outra conexão com outro web client, o que resulta na desconexão da primeira sessão. Esse tipo de priorização da conexão para o mais recente viola o pilar de confidencialidade, permitindo que usuários não autorizados visualizem informações confidenciais.

*Pergunta:* 

Tente simular uma violação do pilar de Integridade.

*Resposta:* 

A simulação de uma conexão bem-sucedida com um broker remoto usando um web client, seguida por outra conexão, desconectando a primeira sessão, viola o pilar de integridade. Isso ocorre porque não há garantia de que as ações realizadas após essa conexão indevida serão refletidas consistentemente, comprometendo assim a integridade dos dados.

*Pergunta:* 

Tente simular uma violação do pilar de Disponibilidade.

*Resposta:* 

A simulação de uma conexão bem-sucedida com um broker remoto usando um web client, seguida por outra conexão que desconecta a primeira sessão, viola o pilar de disponibilidade. Isso ocorre porque a conexão com o broker remoto é priorizada para a conexão mais recente bem-sucedida, comprometendo assim a disponibilidade do serviço para os usuários legítimos.

*Link da evidência dos 3 pilares:*

https://youtu.be/uHVEoTOwI00?si=mDqRcQf84qfTArIf